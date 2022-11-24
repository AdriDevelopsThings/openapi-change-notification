import '../styles/globals.css'
import '../node_modules/bootstrap/dist/css/bootstrap.min.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'

export default function App({ Component, pageProps }: AppProps) {
  return <>
    <Head>
      <title>OpenAPI notifications</title>
      <meta charSet='UTF-8' />
    </Head>
    <div style={{ marginTop: '1rem' }}>
      <Component {...pageProps} />
    </div>
  </>
}
