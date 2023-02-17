import { Heading, Text, HStack, VStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

import Navbar from '../components/Navbar';

export default function Home() {
    const router = useRouter()

    return (
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
    );
}

