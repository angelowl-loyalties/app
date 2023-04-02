import { Button, Container, Divider, HStack, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import { useRouter } from 'next/router';
import { useEffect, useRef, useState } from 'react';
import Loading from '../loading';
import { useToast } from '@chakra-ui/react'

import Navbar from '../../components/Navbar';
import { useSession } from 'next-auth/react';

export default function Banks() {
    const toast = useToast()
    const router = useRouter()
    const [file, setFile] = useState(null);
    const [loading, setLoading] = useState(true)
    const handleFileChange = (event) => {
        console.log("Changed", event.target.files[0])
        setFile(event.target.files[0]);
    }

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        }
    });

    useEffect(() => {
        if (!session) {
            return
        }
        setLoading(false)
    }, [session])

    const submitFile = () => {
        toast({
            title: 'In progress',
            description: "Please hold while we upload your file",
            status: 'info',
            duration: 9000,
            isClosable: true,
        })
        axios.get(`https://itsag1t2.com/publish/presigned`, { headers: { Authorization: session.id } })
            .then((response) => {
                console.log(response)
                axios.put(response.data, file, {
                    headers: {
                        'Content-Type': 'text/csv'
                    }
                }).then((response) => {
                    console.log(response)
                    toast.closeAll()
                    toast({
                        title: 'Success',
                        description: `File uploaded successfully`,
                        status: 'success',
                        duration: 9000,
                        isClosable: true,
                    })
                }).catch((error) => {
                    console.log(error)
                    toast.closeAll()
                    toast({
                        title: 'Error',
                        description: "An error occurred while uploading your file",
                        status: 'error',
                        duration: 9000,
                        isClosable: true,
                    })
                })
            }).catch((error) => {
                console.log(error)
            })

    }

    const seedFile = () => {
        toast({
            title: 'In progress',
            description: "Please hold while we upload your file",
            status: 'info',
            duration: 9000,
            isClosable: true,
        })
        axios.get(`https://itsag1t2.com/user/presigned`, { headers: { Authorization: session.id } })
            .then((response) => {
                console.log(response)
                axios.put(response.data, file, {
                    headers: {
                        'Content-Type': 'text/csv'
                    }
                }).then((response) => {
                    console.log(response)
                    toast.closeAll()
                    toast({
                        title: 'Success',
                        description: `File uploaded successfully`,
                        status: 'success',
                        duration: 9000,
                        isClosable: true,
                    })
                }).catch((error) => {
                    console.log(error)
                    toast.closeAll()
                    toast({
                        title: 'Error',
                        description: "An error occurred while uploading your file",
                        status: 'error',
                        duration: 9000,
                        isClosable: true,
                    })
                })
            }).catch((error) => {
                console.log(error)
            })

    }

    return (
        <>
            {loading ? <Loading /> :
                <Navbar admin>
                    <HStack mb={{ base: 4, lg: 6 }}>
                        <VStack alignItems='start'>
                            <Text textStyle="title">File Upload</Text>
                            <Text textStyle="subtitle" pb={5}>
                                Upload CSV file for processing
                            </Text>
                            <Container>
                                <Text textStyle="head">Bank Transactions</Text>
                                <input type='file' accept=".csv" onChange={handleFileChange} />
                                <Button onClick={submitFile}>Submit</Button>
                            </Container>

                            <Container>
                                <Text textStyle="head">User Accounts</Text>
                                <input type='file' accept=".csv" onChange={handleFileChange} />
                                <Button onClick={seedFile}>Submit</Button>
                            </Container>
                        </VStack>
                    </HStack>
                </Navbar>
            }
        </>
    );
}

