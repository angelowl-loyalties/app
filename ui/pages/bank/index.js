import { useToast, VStack } from '@chakra-ui/react';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect, useRef, useState } from 'react';

import Loading from '../../components/Loading';
import Navbar from '../../components/Navbar';
import Upload from '../../components/UploadComponent';

export default function Banks() {
    const toast = useToast()
    const router = useRouter()
    const [loading, setLoading] = useState(true)
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

    return (
        <>
            {loading ? <Loading /> :
                <Navbar bank>
                    <VStack alignItems="start" w="full" h="full">
                        <Upload toast={toast} session={session} />
                    </VStack>
                </Navbar>
            }
        </>
    );
}

