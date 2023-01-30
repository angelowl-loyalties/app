import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';
import {
    Avatar,
    Box,
    Button,
    Flex,
    FormControl,
    FormHelperText,
    Heading,
    IconButton,
    Input,
    InputGroup,
    InputLeftElement,
    InputRightElement,
    Link,
    Stack,
} from '@chakra-ui/react';
import { getCsrfToken, signIn, useSession } from 'next-auth/react';
import Image from 'next/image';
import { useEffect } from 'react';
import { useRef, useState } from 'react';
import { FaUserAlt, FaLock } from "react-icons/fa";

export default function Login() {
    const [loading, setLoading] = useState(false);
    const email = useRef();
    const password = useRef();
    const [showPassword, setShowPassword] = useState(false);
    const handleShowClick = () => setShowPassword(!showPassword);

    const handleLogin = async () => {
        setLoading(true)
        if (email.current.value && password.current.value) {
            signIn("credentials", {
                email: email.current.value,
                password: password.current.value,
                csrfToken: getCsrfToken(),
                callbackUrl: "/",
            }).then(function (response) {
                // TODO: Remove response
                console.log(response);
            }).catch(function (e) {
                console.log(e)
                setLoading(false)
            });
        }
    };

    useEffect(() => {
        document.title = "Log in | Ascenda"
    }, [])

    return (
        <Flex
            flexDirection="column"
            width="100wh"
            height="100vh"
            justifyContent="center"
            alignItems="center"
            bgImage="url('/bg.jpg')"
            bgPosition="center"
            bgRepeat="no-repeat"
            bgSize="cover"
            bgBlendMode="lighten"
        >
            <Stack
                flexDir="column"
                mb="2"
                justifyContent="center"
                alignItems="center"
                shadow="md"
                bg={'gray.50'}
                rounded={'xl'}
                p={3}
                height="-webkit-fit-content"
                spacing={{ base: 8 }}
                maxW={{ lg: 'md' }}
            >
                <Box pt={8}>
                    <Image priority={true} src="/ascenda.webp" width='0' height='0' sizes="100vw" style={{ width: '150px', height: 'auto' }} alt="ascenda logo" />
                </Box>
                <Box minW={{ base: "90%", md: "100%" }}>
                    <form>
                        <Stack
                            spacing={4}
                            p="1rem"
                        >
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        color="gray.300"><FaUserAlt /></InputLeftElement>
                                    <Input type="email" placeholder="john@doe.com" ref={email} />
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <InputGroup>
                                    <InputLeftElement
                                        color="gray.300"
                                    ><FaLock /></InputLeftElement>
                                    <Input
                                        type={showPassword ? "text" : "password"}
                                        placeholder="Password"
                                        ref={password}
                                    />

                                    <InputRightElement>
                                        <IconButton
                                            color="gray.300"
                                            aria-label='Show password'
                                            size='md'
                                            variant='ghost'
                                            icon={showPassword ? <ViewOffIcon /> : <ViewIcon />}
                                        />
                                    </InputRightElement>
                                </InputGroup>
                                <FormHelperText textAlign="right">
                                    <Link>Forgot your password?</Link>
                                </FormHelperText>
                            </FormControl>
                            <Button
                                borderRadius='lg'
                                type="submit"
                                variant="solid"
                                colorScheme="purple"
                                width="full"
                                onClick={handleLogin}
                            >
                                Login
                            </Button>
                        </Stack>
                    </form>
                </Box>
            </Stack>
        </Flex>
    );
}