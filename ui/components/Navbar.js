import { Box, Divider, Heading, HStack, Stack, Text, VStack } from '@chakra-ui/react';
import { signOut } from 'next-auth/react';
import Head from 'next/head';
import Image from 'next/image';
import NextLink from 'next/link';
import { useRouter } from 'next/router';
import React, { useMemo } from 'react';
import { RiLogoutCircleRLine } from 'react-icons/ri';

import CustomMarquee from './CustomMarquee';
import MainContent from './MainContent';

function Navbar(props) {
    const router = useRouter()

    var routes = [
        {
            'path': '/',
            'tab': 'Dashboard',
            'header': 'Dashboard'
        },
        {
            'path': '/transactions',
            'tab': 'Transaction history',
            'header': 'Transactions'
        },
        {
            'path': '/cards',
            'tab': 'Linked cards',
            'header': 'Cards'
        },
        {
            'path': '/campaigns',
            'tab': 'View campaigns',
            'header': 'Campaigns'
        },
    ]

    return (
        <>
            <Head>
                <title>{routes.find((route) => route.path == router.pathname).header + " | Ascenda"}</title>
            </Head>
            <Stack minH='100vh' direction={{ base: 'column', md: 'row' }}>
                <VStack alignItems="start" backgroundColor="#f5f7f9" w="26%">
                    <Box p={8}>
                        <Image priority={true} src="/ascenda.png" width='0' height='0' sizes="100vw" style={{ width: '150px', height: 'auto' }} alt="ascenda logo" />
                    </Box>
                    <Box flex={1} w="100%">
                        <Stack spacing={6}>
                            <CustomMarquee />
                        </Stack>
                    </Box>
                    <VStack w={'full'} maxW={'lg'} alignItems="start" px={8} pb={9}>
                        <Heading fontWeight="bold" mb={3} mt={2} fontSize="2xl">Manage account</Heading>
                        {routes.map((route) => {
                            return (
                                <Text key={route} fontSize="sm" fontWeight={600} lineHeight="6" textStyle={router.pathname == route.path ? 'navactive' : 'nav'}>
                                    <NextLink href={route.path} verticalalign="middle">
                                        {(router.pathname == route.path ? "> " : "") + route.tab}
                                    </NextLink>
                                </Text>
                            )
                        })}
                        <Divider />
                        <HStack cursor="pointer" onClick={() => signOut()}>
                            <RiLogoutCircleRLine color='red' />
                            <Text fontSize="sm" fontWeight={600} color={'gra6.600'} lineHeight="7" textColor="red" >
                                Log out
                            </Text>
                        </HStack>
                    </VStack>
                </VStack>
                <Box p={8} w="75%" maxH={'98vh'} overflowY="scroll">
                    <MainContent>
                        {props.children}
                    </MainContent>
                </Box>
            </Stack></>
    );
}

export default Navbar;