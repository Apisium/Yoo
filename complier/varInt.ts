import BufferList from 'bl'

export default (buf: BufferList, ...len: Array<number | number[]>) => {
  len.forEach(i => Array.isArray(i)
    ? (write(buf, i.length), i.forEach(n => write(buf, n)))
    : write(buf, i)
  )
}

const write = (buf: BufferList, len: number) => {
  let num = 1
  let i = 1
  while ((num *= 256) - 1 < len) i++
  const b = Buffer.allocUnsafe(i).fill(0xff)
  b[i - 1] = len - (num / 256)
  buf.append(b)
}
