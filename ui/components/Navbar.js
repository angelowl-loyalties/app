import { Alert, AlertIcon, Box, Divider, Heading, HStack, Stack, Text, VStack } from '@chakra-ui/react';
import { signOut } from 'next-auth/react';
import Head from 'next/head';
import Image from 'next/image';
import NextLink from 'next/link';
import { useRouter } from 'next/router';
import React from 'react';
import { RiLogoutCircleRLine } from 'react-icons/ri';
import { BellIcon } from "@chakra-ui/icons"

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
        {
            'path': '/banks',
            'tab': ' ',
            'header': 'SCIS Bank'
        },
    ]

    return (
        <>
            <Head>
                <title>{routes.find((route) => route.path == router.pathname).header + " | Ascenda"}</title>
            </Head>
            <Alert status='info' h={7} display={props.bank ? "inline-flex": "none"} py={0} >
                <AlertIcon w={4} />
                <Text fontSize="xs">Notice: You are currently logged in as a Organisational User from SCIS Bank</Text>
            </Alert>
            <Stack minH='100vh' direction={{ base: 'column', md: 'column', lg: "row" }} >
                <VStack alignItems="start" backgroundColor="#f5f7f9" w="26%" display={{ base: "none", md: "none", lg: "block" }}>
                    <Box p={8}>
                        <Image priority={true} src="/ascenda.webp" width='0' height='0' sizes="100vw" style={{ width: '150px', height: 'auto' }} alt="ascenda logo" />
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
                            <Text fontSize="sm" fontWeight={600} color={'gra6.600'} lineHeight="7" textColor="red" onClick={() => router.push("/login")}>
                                Log out
                            </Text>
                        </HStack>
                    </VStack>
                </VStack>

                <Box
                    position="fixed"
                    mt={0}
                    as="nav"
                    display={{ base: "block", md: "block", lg: "none" }}
                    w="100%"
                    css={{ backgroundColor: "white" }}
                    zIndex={2}
                >
                    <Box
                        display="flex"
                        pb={2}
                        px={{ base: 4, md: 4 }}
                        w="full"
                        wrap="wrap"
                        align="center"
                        css={{ backgroundColor: "rgba(255, 255, 255, .5)", backdropFilter: 'blur(10px)' }}
                    >
                        <Box flex={1} align="left" >
                            <Stack
                                direction={{ base: 'column', md: 'row' }}
                                display={{ base: 'inline-block', md: 'inline-block' }}
                                width={{ base: 'full', md: 'full' }}
                                float="right"
                                flexGrow={1}
                                mt={{ base: 0, md: 0 }}
                            >
                                <Box ml={2} display={{ base: 'inline-block', md: 'inline-block', lg: "none" }}>
                                    <Menu id="navbar-menu">
                                        <MenuButton as={IconButton}
                                            icon={<HamburgerIcon />}
                                            variant="outline"
                                            aria-label="Options" />
                                        <MenuList>
                                            {routes.map((route) => {
                                                return (
                                                    <Text key={route} fontSize="sm" fontWeight={600} lineHeight="6" textStyle={router.pathname == route.path ? 'navactive' : 'nav'}>
                                                        <NextLink href={route.path} verticalalign="middle">
                                                            <MenuItem>{route.tab}</MenuItem>
                                                        </NextLink>
                                                    </Text>
                                                )
                                            })}
                                            <HStack cursor="pointer" onClick={() => signOut()}>
                                                <MenuItem>
                                                    <RiLogoutCircleRLine color='red' />
                                                    <Text fontSize="sm" fontWeight={600} color={'gra6.600'} lineHeight="7" textColor="red" ml={2}>
                                                        Log out
                                                    </Text>
                                                </MenuItem>
                                            </HStack>
                                        </MenuList>
                                    </Menu>
                                </Box>
                            </Stack>
                        </Box>
                        <HStack>
                            <Flex align="center">
                                <NextLink href="/" passHref>
                                    <Image priority={true} src="/small_logo.webp" width='0' height='0' sizes="100vw" style={{ width: "20px", height: 'auto' }} alt="ascenda logo" />
                                </NextLink>
                            </Flex>
                            <Text>

                            </Text>
                        </HStack>
                        <Box flex={1} align="right" >
                            <Stack
                                direction={{ base: 'column', md: 'row' }}
                                display={{ base: 'inline-block', md: 'flex' }}
                                width={{ base: 'full', md: 'auto' }}
                                float="right"
                                flexGrow={1}
                                mt={{ base: 0, md: 0 }}
                            >
                                <Box display={{ base: 'inline-block', md: 'inline-block', lg: "none" }}>
                                    <Menu id="navbar-menu">
                                        <MenuButton as={IconButton}
                                            icon={<BellIcon />}
                                            variant="ghost"
                                            aria-label="Options" />
                                        <MenuList fontSize="xs" textAlign="center">
                                            You have got no messages
                                        </MenuList>
                                    </Menu>
                                </Box>
                            </Stack>


                        </Box>

                    </Box>
                </Box>
                <Box p={{ base: 6, lg: 8 }} w={{ base: "100%", lg: "75%" }} maxH={'98vh'} pt={{ base: "60px", lg: 8 }}>
                    <MainContent>
                        {props.children}
                    </MainContent>
                </Box>
            </Stack>



        </>

    );
}

export default Navbar;