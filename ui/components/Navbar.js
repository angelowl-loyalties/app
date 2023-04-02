import {
    Alert,
    AlertIcon,
    Flex,
    Box,
    Divider,
    Heading,
    HStack,
    IconButton,
    Menu,
    MenuButton,
    MenuItem,
    MenuList,
    Stack,
    Text,
    VStack,
} from "@chakra-ui/react";
import { HamburgerIcon } from "@chakra-ui/icons";
import { signOut } from "next-auth/react";
import Head from "next/head";
import Image from "next/image";
import NextLink from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { useState, useEffect } from "react";
import { RiLogoutCircleRLine } from "react-icons/ri";
import { BellIcon } from "@chakra-ui/icons";

import CustomMarquee from "./CustomMarquee";
import MainContent from "./MainContent";

function Navbar(props) {
    const router = useRouter();
    const [routes, setRoutes] = useState([]);

    var routesBank = [
        {
            path: "/bank",
            tab: "Dashboard",
            header: "Dashboard",
        }
    ];

    var routesAdmin = [
        {
            path: "/admin",
            tab: "Dashboard",
            header: "Dashboard",
        }
    ];

    var routesUser = [
        {
            path: "/",
            tab: "Dashboard",
            header: "Dashboard",
        },
        {
            path: "/transactions",
            tab: "Transaction history",
            header: "Transactions",
        },
        {
            path: "/cards",
            tab: "Linked cards",
            header: "Cards",
        },
        {
            path: "/campaigns",
            tab: "View campaigns",
            header: "Campaigns",
        }
    ];

    useEffect(() => {
        setRoutes(props.bank ? routesBank : props.admin ? routesAdmin : routesUser);
    }, []);

    return (
        <>
            <Head>
                <title>
                    {routes.length != 0 && routes.find((route) => route.path == router.pathname).header +
                        " | Ascenda"}
                </title>
            </Head>
            <Alert
                status={props.bank ? "info" : "error"}
                h={{base: 7, md: 7}}
                display={props.bank | props.admin ? "inline-flex" : "none"}
                py={0}
                zIndex={99}
                position={{base: "fixed", md: "fixed"}}
            >
                    <AlertIcon h="4" />
                    <Text fontSize="xs" textColor={props.bank ? "blue" : "red"}>
                        Notice: You are currently logged in as a {props.bank ? "Organisational User from SCIS Bank" : "Adminstrator"}
                    </Text>
            </Alert>
            <Stack
                direction={{ base: "column", md: "column", lg: "row" }}
                pt={props.user ? 0 : 7}
                minH="fit-content"
                // h={300}
                backgroundColor="#f5f7f9"
            >
                <VStack
                    backgroundColor="#f5f7f9"
                    alignItems="start"
                    w="26%"
                    display={{ base: "none", md: "none", lg: "block" }}
                >
                    <Box p={8}>
                        <img
                            src="https://ik.imagekit.io/alvinowyong/g1t2/ascenda.webp"
                            width="0"
                            height="0"
                            sizes="100vw"
                            style={{ width: "150px", height: "auto" }}
                            alt="ascenda logo"
                        />
                    </Box>
                    <Box flex={1} w="100%">
                        <Stack spacing={6}>
                            <CustomMarquee />
                        </Stack>
                    </Box>
                    <VStack w={"full"} maxW={"lg"} alignItems="start" px={8} pb={9}>
                        <Heading fontWeight="bold" mb={3} mt={2} fontSize="2xl">
                            Manage account
                        </Heading>
                        {routes.map((route) => {
                            return (
                                <Text
                                    key={route.path}
                                    fontSize="sm"
                                    fontWeight={600}
                                    lineHeight="6"
                                    textStyle={
                                        router.pathname == route.path ? "navactive" : "nav"
                                    }
                                >
                                    <NextLink href={route.path} verticalalign="middle">
                                        {(router.pathname == route.path ? "> " : "") + route.tab}
                                    </NextLink>
                                </Text>
                            );
                        })}
                        <Divider />
                        <HStack cursor="pointer" onClick={() => signOut()}>
                            <RiLogoutCircleRLine color="red" />
                            <Text
                                fontSize="sm"
                                fontWeight={600}
                                color={"gra6.600"}
                                lineHeight="7"
                                textColor="red"
                                onClick={() => router.push("/login")}
                            >
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
                    mt={{ base: -2, lg: 0 }}
                        display="flex"
                        pb={2}
                        px={{ base: 4, md: 4 }}
                        w="full"
                        wrap="wrap"
                        align="center"
                        css={{
                            backgroundColor: "rgba(255, 255, 255, .5)",
                            backdropFilter: "blur(10px)",
                        }}
                    >
                        <Box flex={1} align="left">
                            <Stack
                                direction={{ base: "column", md: "row" }}
                                display={{ base: "inline-block", md: "inline-block" }}
                                width={{ base: "full", md: "full" }}
                                float="right"
                                flexGrow={1}
                                pt={{ base: 2, lg: 0 }}
                                mt={{ base: 0, md: 0 }}
                            >
                                <Box
                                    ml={2}
                                    display={{
                                        base: "inline-block",
                                        md: "inline-block",
                                        lg: "none",
                                    }}
                                >
                                    <Menu id="navbar-menu">
                                        <MenuButton
                                            as={IconButton}
                                            icon={<HamburgerIcon />}
                                            variant="outline"
                                            aria-label="Options"
                                        />
                                        <MenuList>
                                            {routes.map((route) => {
                                                return (
                                                    <Text
                                                        key={route.path + "mobile"}
                                                        fontSize="sm"
                                                        fontWeight={600}
                                                        lineHeight="6"
                                                        textStyle={
                                                            router.pathname == route.path
                                                                ? "navactive"
                                                                : "nav"
                                                        }
                                                    >
                                                        <NextLink href={route.path} verticalalign="middle">
                                                            <MenuItem>{route.tab}</MenuItem>
                                                        </NextLink>
                                                    </Text>
                                                );
                                            })}
                                            <HStack cursor="pointer" onClick={() => signOut()}>
                                                <MenuItem>
                                                    <RiLogoutCircleRLine color="red" />
                                                    <Text
                                                        fontSize="sm"
                                                        fontWeight={600}
                                                        color={"gra6.600"}
                                                        lineHeight="7"
                                                        textColor="red"
                                                        ml={2}
                                                    >
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
                                    <img
                                        src="https://ik.imagekit.io/alvinowyong/g1t2/ascenda.webp"
                                        width="0"
                                        height="0"
                                        sizes="100vw"
                                        style={{ width: "80px", height: "auto" }}
                                        alt="logo"
                                    />
                                </NextLink>
                            </Flex>
                        </HStack>
                        <Box flex={1} align="right">
                            <Stack
                                direction={{ base: "column", md: "row" }}
                                display={{ base: "inline-block", md: "flex" }}
                                width={{ base: "full", md: "auto" }}
                                float="right"
                                flexGrow={1}
                                mt={{ base: 0, md: 0 }}
                            >
                                <Box
                                    display={{
                                        base: "inline-block",
                                        md: "inline-block",
                                        lg: "none",
                                    }}
                                    
                                >
                                    <Menu id="navbar-menu" >
                                        <MenuButton
                                            as={IconButton}
                                            icon={<BellIcon />}
                                            variant="ghost"
                                            aria-label="Options"
                                        />
                                        <MenuList fontSize="xs" textAlign="center">
                                            You have got no messages
                                        </MenuList>
                                    </Menu>
                                </Box>
                            </Stack>
                        </Box>
                    </Box>
                </Box>
                <Box
                    backgroundColor="white"
                    p={{ base: 6, lg: 8 }}
                    w={{ base: "100%", lg: "74%" }}
                    overflow="hidden"
                    h="fit-content"
                    minH="100vh"
                    pt={{ base: "60px", lg: 8 }}
                >
                    <MainContent>{props.children}</MainContent>
                </Box>
            </Stack>
        </>
    );
}

export default Navbar;
