import * as ts from 'typescript'
import { readFileSync, createWriteStream } from 'fs'
import writeVarInt from './varInt'
import BufferList = require('bl')

const YOO = Buffer.from('YOO')
const VERSION = Buffer.from([0x01])
const START = [YOO, VERSION]
const CONSTANT_IDENTIFIER = Buffer.from([0x00])
const CONSTANT_CALL = Buffer.from([0x01])
const CONSTANT_STRING = Buffer.from([0x02])
const CONSTANT_MEMBER = Buffer.from([0x03])
const CONSTANT_VARIABLE = Buffer.from([0x04])
const CONSTANT_NULL = Buffer.from([0x05])
const CONSTANT_NUMBER = Buffer.from([0x06])
const CONSTANT_IMPORT = Buffer.from([0x07])
const CONSTANT_ARROW_FUNCTION = Buffer.from([0x08])
const CONSTANT_TRUE = Buffer.from([0x09])
const CONSTANT_FALSE = Buffer.from([0x10])
const CONSTANT_NEW = Buffer.from([0x11])
const CONSTANT_BLOCK = Buffer.from([0x12])
const CONSTANT_REGISTER = Buffer.from([0x13])

// const writeRegister = (buf, i: number) => {
//   buf.append(CONSTANT_REGISTER)
//   writeVarInt(buf, i)
// }
const addRegister = (buf: BufferList, register: boolean[]) => {
  buf.append(CONSTANT_REGISTER)
  const i = register.length
  register[i] = false
  writeVarInt(buf, i)
  return i
}
const writeString = (buf: BufferList, pool: string[], text: string) => {
  buf.append(CONSTANT_STRING)
  const i = pool.indexOf(text)
  writeVarInt(buf, i === -1 ? pool.push(text) - 1 : i)
}
const writeArguments = (buf: BufferList, args: ts.NodeArray<ts.ParameterDeclaration |
  ts.VariableDeclaration>, pool: string[], register: boolean[]) => args.map(arg => [
    expression(arg.name, buf, pool, register),
    expression(arg.initializer, buf, pool, register)
  ])

const writeBoolean = (buf: BufferList, val: boolean) => buf.append(val ? CONSTANT_TRUE : CONSTANT_FALSE)

const expression = (
  n: ts.Node, buf: BufferList, pool: string[], register: boolean[], noAdd = false
): number => {
  if (n && ts.isExpressionStatement(n)) n = n.expression
  let ret: number
  if (!noAdd) ret = addRegister(buf, register)
  if (!n || n.kind === ts.SyntaxKind.NullKeyword || n.kind === ts.SyntaxKind.UndefinedKeyword) {
    if (!noAdd) buf.append(CONSTANT_NULL)
  } else if (ts.isIdentifier(n)) {
    if (!noAdd) {
      buf.append(CONSTANT_IDENTIFIER)
      const i = pool.indexOf(n.text)
      writeVarInt(buf, i === -1 ? pool.push(n.text) - 1 : i)
    }
  } else if (ts.isCallExpression(n)) {
    const i = expression(n.expression, buf, pool, register)
    const indexes = n.arguments.map(na => expression(na, buf, pool, register))
    buf.append(CONSTANT_CALL)
    writeVarInt(buf, i, indexes)
  } else if (ts.isStringLiteral(n)) {
    if (!noAdd) writeString(buf, pool, n.text)
  } else if (ts.isPropertyAccessExpression(n)) {
    const i = expression(n.expression, buf, pool, register)
    buf.append(CONSTANT_MEMBER)
    writeVarInt(buf, i)
    writeString(buf, pool, n.name.text)
  } else if (ts.isElementAccessExpression(n)) {
    const ia = expression(n.expression, buf, pool, register)
    const ib = expression(n.argumentExpression, buf, pool, register)
    buf.append(CONSTANT_MEMBER)
    writeVarInt(buf, ia, ib)
  } else if (ts.isVariableStatement(n)) {
    const indexes = writeArguments(buf, n.declarationList.declarations, pool, register)
    buf.append(CONSTANT_VARIABLE)
    writeVarInt(buf, indexes.length)
    indexes.forEach(([a, b]) => writeVarInt(buf, a, b))
  } else if (ts.isNumericLiteral(n)) {
    if (!noAdd) {
      buf.append(CONSTANT_NUMBER)
      const num = Number.parseFloat(n.text)
      let text = num.toExponential()
      if (text.length > num.toString().length) text = num.toString()
      const buff = Buffer.from(text)
      writeVarInt(buf, buff.length)
      buf.append(text)
    } /*
  } else if (ts.isImportDeclaration(n)) {
    const mod = n.moduleSpecifier
    if (!ts.isStringLiteral(mod)) throw SyntaxError('Is not supported: ' + n.getText())
    buf.append(CONSTANT_IMPORT)
    expression(mod, buf, pool)
    const binds = n.importClause.namedBindings
    writeVarInt(buf, binds.getChildren().length)
    if (binds) {
      binds.forEachChild((c: ts.ImportSpecifier) => {
        expression(c.propertyName || c.name, buf, pool)
        expression(c.name, buf, pool)
      })
    }
  } else if (ts.isArrowFunction(n)) {
    buf.append(CONSTANT_ARROW_FUNCTION)
    writeBoolean(buf, n.modifiers && n.modifiers.length &&
      n.modifiers[0].kind === ts.SyntaxKind.AsyncKeyword)
    writeArguments(buf, n.parameters, pool)
    expression(n.body, buf, pool)
  } else if (ts.isNewExpression(n)) {
    buf.append(CONSTANT_NEW)
    expression(n.expression, buf, pool)
    writeVarInt(buf, n.arguments.length)
    n.arguments.forEach(na => expression(na, buf, pool)) */
  } else if (ts.isSourceFile(n) || ts.isBlock(n)) {
    const arr: number[] = []
    n.forEachChild(na => arr.push(expression(na, buf, pool, register, true)))
    buf.append(CONSTANT_BLOCK)
    writeVarInt(buf, arr.length, arr)
  } else {
    console.log(ts.SyntaxKind[n.kind])
    console.trace()
  }
  return ret
}

const compile = (code: string) => {
  const pool = []
  const register = []
  const buf = new BufferList()

  expression(ts.createSourceFile('a.ts', code, ts.ScriptTarget.Latest, true), buf, pool, register)

  const b = new BufferList()
  b.append(START)
  writeVarInt(b, pool.length)
  pool.forEach(s => {
    const buff = Buffer.from(s, 'UTF-8')
    writeVarInt(b, buff.length)
    b.append(buff)
  })
  b.append(buf)
  return b
}
export default compile
const buffer = compile(readFileSync('test.ts').toString())
buffer.pipe(createWriteStream('a.yoo'))
console.log(buffer.length, buffer.slice())
