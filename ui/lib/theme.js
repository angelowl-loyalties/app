import { extendTheme } from '@chakra-ui/react'
import '@fontsource/inter'
import '@fontsource/public-sans'

const theme = extendTheme({
    textStyles: {
        title: {
            fontWeight: "bold",
            fontSize: {base: 'md', lg: '2xl'},
            fontWeight: {base: 'bold', lg: "semibold"},
        },
        subtitle: {
            lineHeight: '1',
            color: "gray.500",
            fontSize: {base: 'small', lg: 'sm'},
            fontWeight: {base: 500, lg: 600}
        },
        head: {
            fontWeight: "bold",
            mb: {base: 0, lg: 2},
            fontSize: {base: 'sm', lg: 'lg'},
            fontWeight: {base: 'bold', lg: "semibold"},
        },
        tab: {
            fontSize: {base: "xs", lg: "sm"},
        },
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