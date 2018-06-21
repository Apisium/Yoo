declare module 'net/http' {
  export = http
}
declare namespace http {
  export interface ServerOptions {
    host?: string
    port?: number
    handler?: Handler
    idleTimeout?: number
    readTimeout?: number
    writeTimeout?: number
    maxHeaderBytes?: number
    readHeaderTimeout?: number
  }
  export interface Cookie {
    value: string
    path?: string
    domain?: string
    expires?: string
    rawExpires?: number
    maxAge?: number
    secure?: boolean
    httpOnly?: boolean
    raw?: string
  }
  export interface Content extends io.WriteCloser {
    status: number
    method: string
    readonly path: string
    readonly body: io.ReadCloser
    readonly contentLength: number
    readonly cookies: Map<string, Cookie>
    readonly headers: Map<string, string>
    getCookie(key: string): Cookie
    delHeader (key: string): this
    addHeader (key: string, value: string): this
    setHeader (key: string, value: string): this
    basicAuth (name: string, password: string, ok: boolean): this
    addCookie (name: string, value: Cookie | string): this
  }
  export type Handler = (content: Content) => void
  export class Server1 {
    constructor (options: ServerOptions)
    constructor (handler?: Handler, port?: number)

    listen (): this
    listenTLS (cert: string, key: string): this
    close (): this
  }
  export function Server(handler?: Handler, port?: number): Server1
}

declare module 'io' {
  export = io
}
declare namespace io {
  export interface Closer {
    close (bytes: Uint8Array)
  }
  export interface WriteCloser extends Closer {
    write (bytes: Uint8Array): number
  }
  export interface ReadCloser extends Closer {
    read (): Uint8Array
  }
}

// declare class Url {
//   constructor (url: string, base?: string)
//   hash: string
//   host: string
//   hostname: string
//   href: string
//   origin: string
//   password: string
//   pathname: string
//   port: string
//   protocol: string
//   search: string
//   username: string
// }
