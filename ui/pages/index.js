import { Heading, Text } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useEffect } from 'react';

import Navbar from '../components/Navbar';

export default function Home() {
    const router = useRouter()

    return (
        <Navbar>
            <Heading fontWeight="bold" mb={4} fontSize='2xl'>Dashboard</Heading>
            <Text fontSize="sm" fontWeight={600} color={'gray.600'} lineHeight="4">
                Supercharge your everyday credit cards and get rewarded when you spend
            </Text>
        </Navbar>
    );
}