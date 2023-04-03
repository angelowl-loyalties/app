import { Center, Spinner, VStack } from '@chakra-ui/react';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import React, { useEffect } from 'react';


export default function Loading() {
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
        <Center h="100vh" bg="white">
            <VStack>
                <img
                    src="https://ik.imagekit.io/alvinowyong/g1t2/ascenda.webp"
                    width="0"
                    height="0"
                    sizes="100vw"
                    style={{ width: "150px", height: "auto", marginBottom: "30px" }}
                    alt="ascenda logo"
                />
                <Spinner size="md" />
            </VStack>
        </Center>
    );
}

