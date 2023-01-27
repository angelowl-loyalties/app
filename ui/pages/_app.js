import '../styles/globals.css'
import { ChakraProvider } from '@chakra-ui/react'
import theme from "../lib/theme"
import { SessionProvider } from "next-auth/react"


function MyApp({ Component, pageProps: { session, ...pageProps }, }) {
    return (
        <SessionProvider session={session}>
            <ChakraProvider theme={theme}>
                <Component {...pageProps} />
            </ChakraProvider>
        </SessionProvider>
    )
}


export default MyApp
