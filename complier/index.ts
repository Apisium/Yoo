import * as ts from 'typescript'
import { createWriteStream } from 'fs'
import BufferList = require('bl')

const VERSION = Buffer.from([0x01])
const CONSTANT_STRINGS_START = Buffer.from([0x00])
const CONSTANT_STRINGS_END = Buffer.from([0x01])
const CONSTANT_IDENTIFIER = Buffer.from([0x02])
const CONSTANT_CALL = Buffer.from([0x03])
const CONSTANT_STRING = Buffer.from([0x04])
const CONSTANT_MEMBER = Buffer.from([0x05])
const CONSTANT_VARIABLE = Buffer.from([0x06])
const CONSTANT_NULL = Buffer.from([0x07])
const strings = []

const a = new BufferList()
const write = (buf: Buffer) => {
  const length = Buffer.allocUnsafe(8)
  length.writeIntBE(buf.length, 0, 8)
  a.append(length)
  a.append(buf)
  return buf.length + 8
}

const code = 'const text = null; console.log("Hello World!", text)'

const node = ts.createSourceFile(__dirname + '/a.ts', code, ts.ScriptTarget.Latest, true)
const str = (text: string) => {
  const i = strings.indexOf(text)
  return i === -1 ? strings.push(text) - 1 : i
}
const expression = (n: ts.Node) => {
  if (ts.isExpressionStatement(n)) n = n.expression
  if (ts.isIdentifier(n)) {
    a.append(CONSTANT_IDENTIFIER)
    write(Buffer.from(n.text))
  } else if (ts.isCallExpression(n)) {
    const s = a.length
    a.append(CONSTANT_CALL)
    expression(n.expression)
    const length = Buffer.allocUnsafe(2)
    length.writeInt16BE(n.arguments.length, 0)
    a.append(length)
    n.arguments.forEach(expression)
  } else if (ts.isStringLiteral(n)) {
    a.append(CONSTANT_STRING)
    const id = Buffer.allocUnsafe(2)
    id.writeInt16BE(str(n.text), 0)
    a.append(id)
  } else if (ts.isPropertyAccessExpression(n)) {
    a.append(CONSTANT_MEMBER)
    expression(n.expression)
    a.append(CONSTANT_STRING)
    const id = Buffer.allocUnsafe(2)
    id.writeInt16BE(str(n.name.text), 0)
    a.append(id)
  } else if (ts.isElementAccessExpression(n)) {
    a.append(CONSTANT_MEMBER)
    expression(n.expression)
    expression(n.argumentExpression)
  } else if (ts.isVariableStatement(n)) {
    const des = n.declarationList.declarations
    a.append(CONSTANT_VARIABLE)
    const length = Buffer.allocUnsafe(2)
    length.writeInt16BE(des.length, 0)
    a.append(length)
    des.forEach(d => {
      expression(d.name)
      expression(d.initializer)
    })
  } else if (!n || n.kind === ts.SyntaxKind.NullKeyword || n.kind === ts.SyntaxKind.UndefinedKeyword) {
    a.append(CONSTANT_NULL)
  } else if (ts.isInterfaceDeclaration(n)) {
    n.members.forEach(e => {
      if (ts.isPropertySignature(e)) {
        console.log(ts.SyntaxKind[e.type.kind])
      }
    })
  } else console.log(ts.SyntaxKind[n.kind])
}
ts.forEachChild(node, expression)
const b = new BufferList()
b.append(Buffer.from('YOO'))
b.append([VERSION, CONSTANT_STRINGS_START])
const l = Buffer.allocUnsafe(2)
l.writeInt16BE(strings.length, 0)
b.append(l)
strings.forEach((s, i) => {
  const buf = Buffer.from(s, 'UTF-8')
  const length = Buffer.allocUnsafe(8)
  length.writeIntBE(buf.length, 0, 8)
  b.append(length)
  b.append(buf)
})
b.append(CONSTANT_STRINGS_END)
b.append(a)
b.pipe(createWriteStream('a.yoo'))
console.log(b.length, b.slice())
