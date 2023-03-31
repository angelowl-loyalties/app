import { HStack, Text, VStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useState, useEffect } from 'react';
import Navbar from '../components/Navbar';
import Loading from './loading';
import { useSession } from 'next-auth/react';

export default function Home() {
    const [loading, setLoading] = useState(true)
    const router = useRouter()

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        },
    });
    
    useEffect(() => {
        setLoading(false)
    },[])

    return (
        <>
            {loading ? <Loading /> :
                <Navbar>
                    <HStack mb={{ base: 4, lg: 6 }}>
                        <VStack alignItems='start'>
                            <Text textStyle="title">Dashboard</Text>
                            <Text textStyle="subtitle">
                                Supercharge your everyday credit cards and get rewarded when you spend
                            </Text>
                        </VStack>
                    </HStack>
                </Navbar>
            }
        </>
    );
}

