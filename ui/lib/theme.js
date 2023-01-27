import { extendTheme } from '@chakra-ui/react'
import '@fontsource/inter'
import '@fontsource/public-sans'

const theme = extendTheme({
    textStyles: {
        nav: {
            'color': "gray",
            'cursor': "pointer",
            '_hover': {
                'color': '#00a3c4'
            },
        },
        navactive: {
            'cursor': "pointer",
            'color': "#00a3c4",
        }
    },
    fonts: {
        body: `'Inter', sans-serif`,
        heading: `'Public Sans', sans-serif`,
    }
})
export default theme