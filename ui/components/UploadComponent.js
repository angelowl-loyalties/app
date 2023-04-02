import { Box, Button, Container, Stack, Text } from '@chakra-ui/react';
import axios from 'axios';
import { useRouter } from 'next/router';
import { useEffect, useRef, useState } from 'react';

export default function Upload(props) {
    const router = useRouter()
    const [file, setFile] = useState(null);
    const [seedFile, setSeedFile] = useState(null);
    const [loading, setLoading] = useState(true)

    const handleDragOver = (event) => {
        event.preventDefault();
    };
    const handleSeedDragOver = (event) => {
        event.preventDefault();
    };

    const handleDrop = (event) => {
        event.preventDefault();
        console.log("Changed transaction", event.dataTransfer.files[0]);
        setFile(event.dataTransfer.files[0]);
    };
    const handleSeedDrop = (event) => {
        event.preventDefault();
        console.log("Changed seed", event.dataTransfer.files[0]);
        setSeedFile(event.dataTransfer.files[0]);
    };

    const submitFile = (type) => {
        props.toast({
            title: 'In progress',
            description: "Please hold while we upload your file",
            status: 'info',
            duration: 9000,
            isClosable: true,
        })
        axios.get(`https://itsag1t2.com/${type}/presigned`, { headers: { Authorization: props.session.id } })
            .then((response) => {
                var f_file = null
                if (type == "seed") {
                    if (!seedFile) {
                        props.toast.closeAll()
                        props.toast({
                            title: 'No file found',
                            description: "Please upload a file",
                            status: 'warning',
                            duration: 9000,
                            isClosable: true,
                        })
                        return
                    }
                    f_file = seedFile
                } else {
                    if (!file) {
                        props.toast.closeAll()
                        props.toast({
                            title: 'No file found',
                            description: "Please upload a file",
                            status: 'warning',
                            duration: 9000,
                            isClosable: true,
                        })
                        return
                    }
                    f_file = file
                }
                console.log(response)
                axios.put(response.data, f_file, {
                    headers: {
                        'Content-Type': 'text/csv'
                    }
                }).then((response) => {
                    console.log(response)
                    props.toast.closeAll()
                    props.toast({
                        title: 'Success',
                        description: `File uploaded successfully`,
                        status: 'success',
                        duration: 9000,
                        isClosable: true,
                    })
                }).catch((error) => {
                    console.log(error)
                    props.toast.closeAll()
                    props.toast({
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
            <Text textStyle="head" pt={8}>Manage data</Text>
            <Stack w="100%" spacing={{base: 2, md: 10}} direction={{ base: "column", md: "row" }}>
                <Container px={0} w={{ base: "100%", md: "50%" }} ml={0}>
                    <Box
                        border="2px dashed"
                        borderColor="gray.200"
                        p="4"
                        borderRadius="md"
                        onDragOver={handleDragOver}
                        onDrop={handleDrop}
                    >
                        <Text textStyle="head">Bank Transactions</Text>
                        <Box>
                            {file ? (
                                <Box><Text fontSize="small" textColor="gray.500">{file && file.name}</Text></Box>
                            ) : (
                                <Box><Text fontSize="small" textColor="gray.500">Drag and drop a file here</Text></Box>
                            )}
                        </Box>
                        <Button size="sm" mt={3} onClick={() => submitFile("publish")}>Submit</Button>
                    </Box>
                </Container>
                <Container px={0} w={{ base: "100%", md: "50%" }} display={props.admin ? "block" : "none"}  ml={0}>
                    <Box
                        border="2px dashed"
                        borderColor="gray.200"
                        p="4"
                        borderRadius="md"
                        onDragOver={handleSeedDragOver}
                        onDrop={handleSeedDrop}
                    >
                        <Text textStyle="head">Seed User(s)</Text>
                        <Box>
                            {seedFile ? (
                                <Box><Text fontSize="small" textColor="gray.500">{seedFile && seedFile.name}</Text></Box>
                            ) : (
                                <Box><Text fontSize="small" textColor="gray.500">Drag and drop a file here</Text></Box>
                            )}
                        </Box>
                        <Button size="sm" mt={3} onClick={() => submitFile("user")}>Submit</Button>
                    </Box>
                </Container>
            </Stack>
        </>
    );
}

