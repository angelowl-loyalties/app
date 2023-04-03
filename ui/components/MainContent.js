import { chakra, shouldForwardProp } from '@chakra-ui/react';
import { motion } from 'framer-motion';
import { useRouter } from 'next/router';
import React, { useEffect } from 'react';
import { getCsrfToken, signIn, useSession } from "next-auth/react";

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
            console.log(session)
            router.push("/login")
        },
    });
    useEffect(()=>{
        if (session && session.is_new){
            router.push("/changePassword")
        }
    },[session])

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