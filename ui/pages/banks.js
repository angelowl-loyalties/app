import { Button, HStack, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import { useRouter } from 'next/router';
import { useEffect, useRef, useState } from 'react';

import Navbar from '../components/Navbar';

export default function Banks() {
    const router = useRouter()
    const [file, setFile] = useState(null);
    const handleFileChange = (event) => {
        console.log("Changed",event.target.files[0] )
        setFile(event.target.files[0]);
    }

    const submitFile = () => {
        axios.get("https://bkhp6p2ncthmrlv6dftdtyzzxm0xrvfr.lambda-url.ap-southeast-1.on.aws/")
            .then((response) => {
                console.log(response)
                axios.put(response.data, file, {
                    headers: {
                        'Content-Type': 'text/csv'
                    }
                }).then((response) => {
                    console.log(response)
                }).catch((error) => {
                    console.log(error)
                })
            }).catch((error) => {
                console.log(error)
            })

    }

    return (
        <Navbar bank>
            <HStack mb={{ base: 4, lg: 6 }}>
                <VStack alignItems='start'>
                    <Text textStyle="title">Transactions Upload</Text>
                    <Text textStyle="subtitle">
                        Upload CSV file for processing
                    </Text>
                    <input type='file' accept=".csv" onChange={handleFileChange}/>
                    <Button onClick={submitFile}>Submit</Button>
                </VStack>
            </HStack>
        </Navbar>
    );
}

