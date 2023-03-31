import { Search2Icon } from '@chakra-ui/icons';
import {
    Box,
    HStack,
    Input,
    InputGroup,
    InputLeftElement,
    Select,
    Spacer,
    Tab,
    TabList,
    TabPanels,
    Tabs,
    Text,
    VStack,
} from '@chakra-ui/react';
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { FaCcAmex, FaCcDiscover, FaCcMastercard, FaCcVisa, FaCreditCard } from 'react-icons/fa';

import Navbar from '../components/Navbar';
import RewardPanel from '../components/RewardPanel';
import Loading from './loading';

export default function Transactions() {
    const router = useRouter()
    const [cards, setCards] = useState({})
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
        axios.get(`https://itsag1t2.com/user/${session.userId}`, { headers: { Authorization: session.id } })
            .then((response) => {
                console.log(response.data.data.CreditCards)
                setCards(response.data.data.CreditCards)
                setLoading(false)
            })


    }, [session])


    const handleSearch = (e) => {
        console.log("Hello, not implemented 🚀")
    }

    return (
        <>
            {loading ? <Loading /> :
                <Navbar user>
                    <VStack alignItems='start' w="full">
                        <HStack mb={{ base: 4, lg: 6 }}>
                            <VStack alignItems='start'>
                                <Text textStyle="title">Transactions and points history</Text>
                                <Text textStyle="subtitle">
                                    Supercharge your credit cards and get rewarded when you spend
                                </Text>
                            </VStack>
                        </HStack>
                        <Tabs variant='solid-rounded' colorScheme="purple" w="full">
                            <HStack>
                                <Select
                                    w="25%"
                                    fontSize="small"
                                    display={{ base: "inline-block", lg: "none" }}
                                    placeholder='All'>
                                    {cards.map((card) => {
                                        const num = card.card_pan.slice(-4)
                                        switch (card.card_pan.charAt(0)) {
                                            case '3':
                                                return <option fontSize='sm' style={{ borderRadius: 'lg' }} key={num}>(AMEX) {num}</option>
                                            case '4':
                                                return <option fontSize='sm' style={{ borderRadius: 'lg' }} key={num}>(VISA) {num}</option>
                                            case '5':
                                                return <option fontSize='sm' style={{ borderRadius: 'lg' }} key={num}>(MCC) {num}</option>
                                            case '6':
                                                return <option fontSize='sm' style={{ borderRadius: 'lg' }} key={num}>(DISC) {num}</option>
                                            default:
                                                return <option fontSize='sm' style={{ borderRadius: 'lg' }} key={num}>(OTHERS){num}</option>
                                        }
                                    })}
                                </Select>
                                <Box p={2} bgColor="gray.100" borderRadius="xl" display={{ base: "none", lg: "inline-block" }}>
                                    <TabList>
                                        {cards.map((card) => {
                                            const num = card.card_pan.slice(-4)
                                            switch (card.card_pan.charAt(0)) {
                                                case '3':
                                                    return <Tab fontSize='sm' borderRadius='lg' key={num}><HStack><FaCcAmex size={23} /><Text>{num}</Text></HStack></Tab>
                                                case '4':
                                                    return <Tab fontSize='sm' borderRadius='lg' key={num}><HStack><FaCcVisa size={23} /><Text>{num}</Text></HStack></Tab>
                                                case '5':
                                                    return <Tab fontSize='sm' borderRadius='lg' key={num}><HStack><FaCcMastercard size={23} /><Text>{num}</Text></HStack></Tab>
                                                case '6':
                                                    return <Tab fontSize='sm' borderRadius='lg' key={num}><HStack><FaCcDiscover size={23} /><Text>{num}</Text></HStack></Tab>
                                                default:
                                                    return <Tab fontSize='sm' borderRadius='lg' key={num}><HStack><FaCreditCard size={23} /><Text>{num}</Text></HStack></Tab>
                                            }
                                        })}
                                    </TabList>
                                </Box>
                                <Spacer />size={20}
                                <InputGroup w="30%">
                                    <InputLeftElement
                                        pointerEvents='none'><Search2Icon color='gray.300' /></InputLeftElement>
                                    <Input type='text' placeholder='Search' fontSize='sm' onChange={handleSearch} />
                                </InputGroup>
                            </HStack>

                            <TabPanels>
                                {cards.map((card) => {
                                    return (
                                        <RewardPanel card={card} key={card.id} />
                                    )
                                })}
                            </TabPanels>
                        </Tabs>
                    </VStack>
                </Navbar>}
        </>
    );
}


