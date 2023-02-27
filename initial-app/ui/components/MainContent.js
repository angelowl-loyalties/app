import { motion } from 'framer-motion'
import { chakra, shouldForwardProp } from '@chakra-ui/react'
import React from 'react'
import { getCsrfToken, signIn, useSession } from 'next-auth/react';
import { useRouter } from 'next/router';

const StyledDiv = chakra(motion.div, {
    shouldForwardProp: prop => {
        return shouldForwardProp(prop) || prop === 'transition'
    }
})

const MainContent = ({ children }) => {
    const router = useRouter()
    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            // router.push('/login')
        },
    })

    return (
        <StyledDiv
            initial={{ y: 10, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ duration: 0.5 }}
            mb={6}
        >
            {children}
        </StyledDiv>
    )
}

export default MainContent