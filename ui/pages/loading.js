import { Center, Spinner, VStack } from '@chakra-ui/react';
import { useRouter } from 'next/router';


export default function Loading() {
    const router = useRouter()

    return (
        <Center h="100vh" bg="white">
            <VStack>
                <img
                    priority={true}
                    src="https://ik.imagekit.io/alvinowyong/g1t2/ascenda.webp"
                    width="0"
                    height="0"
                    sizes="100vw"
                    style={{ width: "150px", height: "auto", marginBottom: "30px" }}
                    alt="ascenda logo"
                />
                <Spinner size="md" />
            </VStack>
        </Center>
    );
}
