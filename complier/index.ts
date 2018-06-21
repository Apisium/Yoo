import * as ts from 'typescript'
import { readFileSync, createWriteStream } from 'fs'
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

const writeIdentifier = (buf: BufferList, pool: string[], text: string) => {
  buf.append(CONSTANT_IDENTIFIER)
  const i = pool.indexOf(text)
  writeLength(buf, i === -1 ? pool.push(text) - 1 : i)
}
const writeString = (buf: BufferList, pool: string[], text: string) => {
  buf.append(CONSTANT_STRING)
  const i = pool.indexOf(text)
  writeLength(buf, i === -1 ? pool.push(text) - 1 : i)
}
const writeLength = (buf: BufferList, len: number) => {
  const length = Buffer.allocUnsafe(2)
  length.writeInt16BE(len, 0)
  buf.append(length)
}
const writeArguments = (buf: BufferList, args: ts.NodeArray<ts.ParameterDeclaration |
  ts.VariableDeclaration>, pool: string[]) => {
  writeLength(buf, args.length)
  args.forEach(arg => {
    expression(arg.name, buf, pool)
    expression(arg.initializer, buf, pool)
  })
}
const writeBoolean = (buf: BufferList, val: boolean) => buf.append(val ? CONSTANT_TRUE : CONSTANT_FALSE)

const expression = (n: ts.Node, buf: BufferList, pool: string[]) => {
  if (n && ts.isExpressionStatement(n)) n = n.expression
  if (!n || n.kind === ts.SyntaxKind.NullKeyword || n.kind === ts.SyntaxKind.UndefinedKeyword) {
    buf.append(CONSTANT_NULL)
  } else if (ts.isIdentifier(n)) {
    writeIdentifier(buf, pool, n.text)
  } else if (ts.isCallExpression(n)) {
    buf.append(CONSTANT_CALL)
    expression(n.expression, buf, pool)
    writeLength(buf, n.arguments.length)
    n.arguments.forEach(na => expression(na, buf, pool))
  } else if (ts.isStringLiteral(n)) {
    writeString(buf, pool, n.text)
  } else if (ts.isPropertyAccessExpression(n)) {
    buf.append(CONSTANT_MEMBER)
    expression(n.expression, buf, pool)
    writeString(buf, pool, n.name.text)
  } else if (ts.isElementAccessExpression(n)) {
    buf.append(CONSTANT_MEMBER)
    expression(n.expression, buf, pool)
    expression(n.argumentExpression, buf, pool)
  } else if (ts.isVariableStatement(n)) {
    if (n.getFirstToken().kind === ts.SyntaxKind.VarKeyword) throw new SyntaxError('Is not supported: Var keyword.')
    buf.append(CONSTANT_VARIABLE)
    writeArguments(buf, n.declarationList.declarations, pool)
  } else if (ts.isInterfaceDeclaration(n)) {
    n.members.forEach(e => {
      console.log(ts.SyntaxKind[e.kind])
      if (ts.isPropertySignature(e)) {
      }
    })
  } else if (ts.isNumericLiteral(n)) {
    buf.append(CONSTANT_NUMBER)
    const num = Number.parseFloat(n.text)
    let text = num.toExponential()
    if (text.length > num.toString().length) text = num.toString()
    const buff = Buffer.from(text)
    writeLength(buf, buff.length)
    buf.append(text)
  } else if (ts.isImportDeclaration(n)) {
    const mod = n.moduleSpecifier
    if (!ts.isStringLiteral(mod)) throw SyntaxError('Is not supported: ' + n.getText())
    buf.append(CONSTANT_IMPORT)
    expression(mod, buf, pool)
    const binds = n.importClause.namedBindings
    let i = 0
    if (binds) binds.forEachChild(() => void i++)
    writeLength(buf, i)
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
    const body = n.body
    let i = 0
    if (ts.isBlock(body)) body.forEachChild(() => void i++)
    else i = 1
    writeLength(buf, i)
    body.forEachChild(na => expression(na, buf, pool))
  } else if (ts.isNewExpression(n)) {
    buf.append(CONSTANT_NEW)
    expression(n.expression, buf, pool)
    writeLength(buf, n.arguments.length)
    n.arguments.forEach(na => expression(na, buf, pool))
  } else console.log(ts.SyntaxKind[n.kind])
}

const compile = (code: string) => {
  const node = ts.createSourceFile(__dirname + '/a.ts', code, ts.ScriptTarget.Latest, true)
  let i = 0
  const pool = []
  const buf = new BufferList()
  node.forEachChild(() => void i++)
  node.forEachChild(n => expression(n, buf, pool))

  const b = new BufferList()
  b.append(START)
  writeLength(b, pool.length)
  pool.forEach(s => {
    const buff = Buffer.from(s, 'UTF-8')
    writeLength(b, buff.length)
    b.append(buff)
  })
  writeLength(b, i)
  b.append(buf)
  return b
}
export default compile
const buffer = compile(readFileSync('test.ts').toString())
buffer.pipe(createWriteStream('a.yoo'))
console.log(buffer.length, buffer.slice())
