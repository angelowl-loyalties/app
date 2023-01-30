import { extendTheme } from '@chakra-ui/react'
import '@fontsource/inter'
import '@fontsource/public-sans'

const theme = extendTheme({
    textStyles: {
        nav: {
            'color': "gray",
            'cursor': "pointer",
            '_hover': {
                'color': '#7c4cc4'
            },
        },
        navactive: {
            'cursor': "pointer",
            'color': "#7c4cc4",
        }
    },
    fonts: {
        body: `'Inter', sans-serif`,
        heading: `'Public Sans', sans-serif`,
    }
})
export default theme