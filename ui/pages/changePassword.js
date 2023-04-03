import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';
import {
    Box,
    Button,
    Flex,
    FormControl,
    IconButton,
    Input,
    InputGroup,
    InputLeftElement,
    InputRightElement,
    Stack,
    useToast,
} from '@chakra-ui/react';
import axios from 'axios';
import { getCsrfToken, signIn, useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useRef, useState } from 'react';
import { FaLock, FaUserAlt } from 'react-icons/fa';

import Loading from '../components/Loading';

export default function ChangePassword(props) {
    const router = useRouter();
    const [loading, setLoading] = useState(false);
    const email = useRef();
    const password = useRef();
    const oldPassword = useRef();
    const confirmPassword = useRef();
    const [showPassword, setShowPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);
    const handleShowClick = (confirm) => {
        if (confirm) {
            setShowConfirmPassword(!showConfirmPassword);
        } else {
            setShowPassword(!showPassword);
        }
    };
    const { data: session, status } = useSession({
        required: false,
    });
    const [emailValue, setEmailValue] = useState("");
    const toast = useToast();
    const handleChangePassword = async (e) => {
        setLoading(true);
        e.preventDefault();
        const passwordValue = password.current.value
        if (confirmPassword.current.value === password.current.value) {
            axios.post(
                `https://itsag1t2.com/auth/password/`,
                {
                    email: emailValue,
                    password: password.current.value,
                    confirm_password: confirmPassword.current.value,
                    old_password: oldPassword.current.value,
                },
                { headers: { Authorization: session.id } }
            ).then((response) => {
                signIn("credentials", {
                    email: emailValue,
                    password: passwordValue,
                    csrfToken: getCsrfToken(),
                    // callbackUrl: "/",
                    redirect: false,
                }).then(({ ok, error }) => {
                    if (ok) {
                        router.push("/");
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
            });
        } else {
            setLoading(false);
            toast({
                title: "Mismatched Password",
                description: "Passwords don't match",
                status: "error",
            });
        }
    };

    useEffect(() => {
        if (status === "loading") {
            return;
        }
        if (!session) {
            router.push("/login");
        } else {
        }
        if (typeof window !== "undefined") {
            setEmailValue(sessionStorage.getItem("user_email"));
        }
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
                                                ref={email}
                                                disabled
                                                value={emailValue} />
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
                                                placeholder="Old Password"
                                                ref={oldPassword} />
                                            <InputRightElement>
                                                <IconButton
                                                    color="gray.300"
                                                    aria-label="Show password"
                                                    size="md"
                                                    variant="ghost"
                                                    icon={showPassword ? <ViewOffIcon /> : <ViewIcon />}
                                                    onClick={() => handleShowClick(false)} />
                                            </InputRightElement>
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
                                                    onClick={() => handleShowClick(false)} />
                                            </InputRightElement>
                                        </InputGroup>
                                    </FormControl>
                                    <FormControl>
                                        <InputGroup size="md">
                                            <InputLeftElement color="gray.300">
                                                <FaLock />
                                            </InputLeftElement>
                                            <Input
                                                fontSize="sm"
                                                type={showConfirmPassword ? "text" : "password"}
                                                placeholder="Confirm Password"
                                                ref={confirmPassword} />
                                            <InputRightElement>
                                                <IconButton
                                                    color="gray.300"
                                                    aria-label="Show password"
                                                    size="md"
                                                    variant="ghost"
                                                    icon={showConfirmPassword ? <ViewOffIcon /> : <ViewIcon />}
                                                    onClick={() => handleShowClick(true)} />
                                            </InputRightElement>
                                        </InputGroup>
                                    </FormControl>
                                    <Button
                                        borderRadius="lg"
                                        type="submit"
                                        variant="solid"
                                        colorScheme="purple"
                                        width="full"
                                        onClick={handleChangePassword}>
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
