import * as ts from 'typescript'
import { createWriteStream } from 'fs'
import BufferList = require('bl')

const VERSION = Buffer.from([0x01])
const CONSTANT_STRINGS_START = Buffer.from([0x10])
const CONSTANT_STRINGS_END = Buffer.from([0x11])
const CONSTANT_IDENTIFIER = Buffer.from([0x12])
const CONSTANT_CALL = Buffer.from([0x13])
const CONSTANT_STRING = Buffer.from([0x14])
const strings = []

let a = new BufferList()
const write = (buf: Buffer) => {
  const length = Buffer.allocUnsafe(8)
  length.writeIntBE(buf.length, 0, 8)
  a.append(length)
  a.append(buf)
  return buf.length + 8
}

const node = ts.createSourceFile(__dirname + '/a.ts', 'log("Hello World!"); log("uuuu")',
  ts.ScriptTarget.ESNext, true)
const str = (text: string) => {
  const i = strings.indexOf(text)
  return i === -1 ? strings.push(text) - 1 : i
}
const express = (n: ts.Node) => {
  if (ts.isExpressionStatement(n)) n = n.expression
  if (ts.isIdentifier(n)) {
    a.append(CONSTANT_IDENTIFIER)
    return 1 + write(Buffer.from(n.text))
  } else if (ts.isCallExpression(n)) {
    const s = a.length
    a.append(CONSTANT_CALL)
    const aa = a
    a = new BufferList()
    let i = express(n.expression) + 2
    let length = Buffer.allocUnsafe(2)
    length.writeInt16BE(n.arguments.length, 0)
    a.append(length)
    n.arguments.forEach(e => i += express(e))
    length = Buffer.allocUnsafe(8)
    length.writeIntBE(i, 0, 8)
    aa.append(length)
    aa.append(a)
    a = aa
    return i + 1
  } else if (ts.isStringLiteral(n)) {
    a.append(CONSTANT_STRING)
    const id = Buffer.allocUnsafe(2)
    id.writeInt16BE(str(n.text), 0)
    a.append(id)
    return 3
  } else if (ts.isFunctionDeclaration(n)) {

  }
  return 0
}
node.forEachChild(e => void express(e))
const b = new BufferList()
b.append(Buffer.from('YOOOO'))
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
console.log(b.slice())
