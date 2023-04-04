import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';
import {
    Box,
    Button,
    Flex,
    FormControl,
    FormHelperText,
    IconButton,
    Input,
    InputGroup,
    InputLeftElement,
    InputRightElement,
    Link,
    Stack,
    Text,
    useToast,
} from '@chakra-ui/react';
import { getCsrfToken, signIn, useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useRef, useState } from 'react';
import { FaLock, FaUserAlt } from 'react-icons/fa';

import Loading from '../../components/Loading';

export default function Login() {
    const router = useRouter();
    const [loading, setLoading] = useState(false);
    const email = useRef();
    const password = useRef();
    const [showPassword, setShowPassword] = useState(false);
    const handleShowClick = () => setShowPassword(!showPassword);
    const { data: session, status } = useSession({
        required: false,
    });
    const [errorMessage, setErrorMessage] = useState("");
    const toast = useToast()
    const handleLogin = async (e) => {
        e.preventDefault();
        if (!email.current.value || !password.current.value) {
            setErrorMessage("Email or password cannot be empty");
            return;
        }
        setLoading(true);
        if (email.current.value && password.current.value) {
            signIn("credentials", {
                email: email.current.value,
                password: password.current.value,
                csrfToken: getCsrfToken(),
                redirect: false,
            }).then(({ ok, error }) => {
                if (ok) {
                    router.push("/admin");
                } else {
                    setLoading(false);
                    toast({
                        title: "Invalid Credentials",
                        description: "Incorrect email or password",
                        status: "error",
                    });
                    console.log(error);
                }
            });
        }
    };

    useEffect(() => {
        if (session) {
            router.push("/admin");
        }
        document.title = "Log in | Ascenda";
    }, [session]);

    return (
        <>
            {loading ? <Loading /> :
                <Flex
                flexDirection="column"
                width="100wh"
                height="100vh"
                justifyContent="center"
                alignItems="center"
                bgImage="https://ik.imagekit.io/alvinowyong/g1t2/bg.webp?tr=w-1920,h-1080"
                bgPosition="center"
                bgRepeat="no-repeat"
                bgSize="cover"
                bgBlendMode="lighten">
                <Stack
                    flexDir="column"
                    mb="2"
                    justifyContent="center"
                    alignItems="center"
                    shadow="md"
                    bg={"gray.50"}
                    rounded={"xl"}
                    p={5}
                    height="-webkit-fit-content"
                    spacing={{ base: 8 }}
                    maxW={{ lg: "md" }}>
                    <Box pt={8}>
                        <img
                            src="https://ik.imagekit.io/alvinowyong/g1t2/ascenda.webp"
                            width="0"
                            height="0"
                            sizes="100vw"
                            style={{ width: "150px", height: "auto" }}
                            alt="logo" />
                    </Box>
                    <Box minW={{ base: "90%", md: "100%" }}>
                        <form>
                            <Stack spacing={4} p="1rem">
                                <FormControl>
                                    <InputGroup size="md">
                                        <InputLeftElement color="gray.300">
                                            <FaUserAlt />
                                        </InputLeftElement>
                                        <Input
                                            fontSize="sm"
                                            type="email"
                                            placeholder="john@doe.com"
                                            ref={email} />
                                    </InputGroup>
                                </FormControl>
                                <FormControl>
                                    <InputGroup size="md">
                                        <InputLeftElement color="gray.300">
                                            <FaLock />
                                        </InputLeftElement>
                                        <Input
                                            fontSize="sm"
                                            type={showPassword ? "text" : "password"}
                                            placeholder="Password"
                                            ref={password} />
                                        <InputRightElement>
                                            <IconButton
                                                color="gray.300"
                                                aria-label="Show password"
                                                size="md"
                                                variant="ghost"
                                                icon={showPassword ? <ViewOffIcon /> : <ViewIcon />}
                                                onClick={handleShowClick} />
                                        </InputRightElement>
                                    </InputGroup>
                                    <FormHelperText textAlign="right">
                                        <Link>Forgot your password?</Link>
                                    </FormHelperText>
                                </FormControl>
                                <Button
                                    borderRadius="lg"
                                    type="submit"
                                    variant="solid"
                                    colorScheme="purple"
                                    width="full"
                                    onClick={handleLogin}>
                                    Login
                                </Button>
                            </Stack>
                        </form>
                    </Box>
                </Stack>
            </Flex>
            }
        </>
    );
}
