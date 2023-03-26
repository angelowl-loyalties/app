import { Search2Icon } from '@chakra-ui/icons';
import {
    Accordion,
    AccordionButton,
    AccordionIcon,
    Select,
    AccordionItem,
    AccordionPanel,
    Badge,
    Box,
    Divider,
    Heading,
    HStack,
    Input,
    InputGroup,
    InputLeftElement,
    Spacer,
    Tab,
    TabList,
    TabPanel,
    TabPanels,
    Tabs,
    Text,
    VStack,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { CiDeliveryTruck, CiForkAndKnife, CiRoute, CiShoppingBasket, CiWallet } from 'react-icons/ci';
import { FaCcAmex, FaCcDiscover, FaCcMastercard, FaCcVisa, FaCreditCard } from 'react-icons/fa';

import Navbar from '../components/Navbar';
import axios from 'axios';

export default function Transactions() {
    const router = useRouter()
    const [cards, setCards] = useState([])
    const [filteredTransactions, setFilteredTransactions] = useState([])


    useEffect(() => {
        // TODO: Set user id from auth
        axios.get(`https://itsag1t2.com/user/dd03a9d8-a741-4b72-92bc-46096e3d3b00`, { headers: { Authorization: 'Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImlzcyI6ImFuZ2Vsb3dsLmNvbSIsInN1YiI6ImRkMDNhOWQ4LWE3NDEtNGI3Mi05MmJjLTQ2MDk2ZTNkM2IwMCIsImF1ZCI6WyJhcGkuYW5nZWxvd2wuY29tIl0sImV4cCI6MTY3OTkwNDE1MCwibmJmIjoxNjc5ODE3NzUwLCJpYXQiOjE2Nzk4MTc3NTB9.F9aCxhsbkhWMjlZc0iHdhel3Dhirt7Qf06QW_-J_2h3flnv_vxL1Quk8QUv6YqUawB3VMq7s_fIwnDSihuYLGyRbyY43p1ZF9xFR7m3-mGEd_rpgWrQuVoKk9ZNcBu-CXi2RvCSmXN8U93tpAnSOJK0yIZLw5uJNkkZNF6T5_THjkAx4EI9hvPqXdfsrkxyUA-RqxQb9cSA2ktDnyTq1OOmcti6Ylp5WC7ZUzJsuyhEkDEAqYLDZOfjZRKwauTC4-8VZP6NgTDkEgkAOa-mTeG3jkzO7A2x36nabVzhdeyQXpJ1K1bLFLmltgTF-Eq1B7RL5EtypaLFz0leu0wYoiiM0G49--4KEvdAMi-xDLIcUA_DA0yE6IsIfTXzHuwpp_sbom9Tuqifh3nFgumS7w_3fN1UTGWjVDBCsVQLHQrS_3QfIybDC3LVfAL0qDaQCeb-dcmT9_8qU11Yb7jAtp6KoIAnOo1HaBaOBGXEe3VWJsQKgJLoEQcq_tUWj7IbAuQbxXUPTDlrH1pp3yw70yJ7qzfPzP8NsLckMAe5SCDDYJkqh1TizEgwXcVGBCUUQTi9ievrxkVgMtKslcGBdmlKNcvQCO6HCn7XeKP6k_1IBOd4raIU2zkkUtv6ebnBElSH6uDONUECwktVq7FXVmlV0EVVtK86vhQXIqdeenOI' } })
            .then((response) => {
                console.log(response.data.data.CreditCards)
                setCards(response.data.data.CreditCards.map((card) => card.card_pan))

                response.data.data.CreditCards.map((card) => {
                    axios.get(`https://itsag1t2.com/reward/${card.id}`, { headers: { Authorization: 'Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsImlzcyI6ImFuZ2Vsb3dsLmNvbSIsInN1YiI6ImRkMDNhOWQ4LWE3NDEtNGI3Mi05MmJjLTQ2MDk2ZTNkM2IwMCIsImF1ZCI6WyJhcGkuYW5nZWxvd2wuY29tIl0sImV4cCI6MTY3OTkwNDE1MCwibmJmIjoxNjc5ODE3NzUwLCJpYXQiOjE2Nzk4MTc3NTB9.F9aCxhsbkhWMjlZc0iHdhel3Dhirt7Qf06QW_-J_2h3flnv_vxL1Quk8QUv6YqUawB3VMq7s_fIwnDSihuYLGyRbyY43p1ZF9xFR7m3-mGEd_rpgWrQuVoKk9ZNcBu-CXi2RvCSmXN8U93tpAnSOJK0yIZLw5uJNkkZNF6T5_THjkAx4EI9hvPqXdfsrkxyUA-RqxQb9cSA2ktDnyTq1OOmcti6Ylp5WC7ZUzJsuyhEkDEAqYLDZOfjZRKwauTC4-8VZP6NgTDkEgkAOa-mTeG3jkzO7A2x36nabVzhdeyQXpJ1K1bLFLmltgTF-Eq1B7RL5EtypaLFz0leu0wYoiiM0G49--4KEvdAMi-xDLIcUA_DA0yE6IsIfTXzHuwpp_sbom9Tuqifh3nFgumS7w_3fN1UTGWjVDBCsVQLHQrS_3QfIybDC3LVfAL0qDaQCeb-dcmT9_8qU11Yb7jAtp6KoIAnOo1HaBaOBGXEe3VWJsQKgJLoEQcq_tUWj7IbAuQbxXUPTDlrH1pp3yw70yJ7qzfPzP8NsLckMAe5SCDDYJkqh1TizEgwXcVGBCUUQTi9ievrxkVgMtKslcGBdmlKNcvQCO6HCn7XeKP6k_1IBOd4raIU2zkkUtv6ebnBElSH6uDONUECwktVq7FXVmlV0EVVtK86vhQXIqdeenOI' } })
                    .then((response) => {
                        console.log(response.data.data)
                        const truncatedTransactions = response.data.data.map(transaction => ({
                            ...transaction,
                            amount: parseFloat(transaction.amount.toFixed(2))
                          }));
                        setFilteredTransactions(prevTransactions => [...prevTransactions, ...truncatedTransactions])
                    }).catch((error) => {
                        console.log(error)
                    })
                })
            })


    }, [])

    const handleSearch = (e) => {
        const query = e.target.value
        setFilteredTransactions(transactions.filter())
    }

    return (
        <Navbar>
            <VStack alignItems='start' w="full">
                <HStack mb={{ base: 4, lg: 6 }}>
                    <VStack alignItems='start'>
                        <Text textStyle="title">Transactions and points history</Text>
                        <Text textStyle="subtitle">
                            Supercharge your credit cards and get rewarded when you spend
                        </Text>
                    </VStack>
                </HStack>
                {/* TODO: Enable lazy loading */}
                <Tabs variant='solid-rounded' colorScheme="purple" w="full">
                    <HStack>
                        <Select
                            w="25%"
                            fontSize="small"
                            display={{ base: "inline-block", lg: "none" }}
                            placeholder='All'>
                            {cards.map((card) => {
                                const num = card.slice(-4)
                                switch (card.charAt(0)) {
                                    case '3':
                                        return <option fontSize='sm' style={{borderRadius: 'lg'}} key={num}>(AMEX) {num}</option>
                                    case '4':
                                        return <option fontSize='sm' style={{borderRadius: 'lg'}} key={num}>(VISA) {num}</option>
                                    case '5':
                                        return <option fontSize='sm' style={{borderRadius: 'lg'}} key={num}>(MCC) {num}</option>
                                    case '6':
                                        return <option fontSize='sm' style={{borderRadius: 'lg'}} key={num}>(DISC) {num}</option>
                                    default:
                                        return <option fontSize='sm' style={{borderRadius: 'lg'}} key={num}>(OTHERS){num}</option>
                                }
                            })}
                        </Select>
                        <Box p={2} bgColor="gray.100" borderRadius="xl" display={{ base: "none", lg: "inline-block" }}>
                            <TabList>
                                <Tab fontSize='sm' borderRadius='lg'><Text mx={1}>All</Text></Tab>
                                {cards.map((card) => {
                                    const num = card.slice(-4)
                                    switch (card.charAt(0)) {
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
                        <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            exit={{ opacity: 0 }} key="all">
                            <Accordion allowMultiple w="full">
                                {filteredTransactions.map((transaction) => {
                                    return (
                                        <AccordionItem key={transaction.id}>
                                            <AccordionButton p={{ base: 1, lg: 2 }}>
                                                <Box as="span" flex='1' textAlign='left'>
                                                    <HStack pr={{ base: 0, lg: 8 }}>
                                                        <Text display={{ base: "none", md: "block" }}>{(transaction.mcc == '5499' | transaction.mcc == '5411') ? <CiShoppingBasket size={20} color='gray' /> : (transaction.mcc == '5541' | transaction.mcc == '5542') ? <CiDeliveryTruck size={20} color='gray' /> : (transaction.mcc == '5499') ? <CiForkAndKnife size={20} color='gray' /> : (transaction.mcc == '4121' | transaction.mcc == '5734' | transaction.mcc == '6540' | transaction.mcc == '4111') ? <CiRoute size={20} color='gray' /> : (transaction.mcc == '5999' | transaction.mcc == '5964' | transaction.mcc == '5691' | transaction.mcc == '5311' | transaction.mcc == '5411' | transaction.mcc == '5399' | transaction.mcc == '5311') ? <CiShoppingBasket size={20} color='gray' /> : <CiWallet size={20} color='gray' />}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 14, lg: 20 }}>{transaction.transaction_date}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 100, lg: 200 }}>{transaction.merchant}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 110, lg: 200 }}>Shopee 11.11 Sale</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 100, lg: 100 }}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                        <Spacer />
                                                        <Badge fontSize='xs' colorScheme="green" py={1} px={2} >+100 Miles</Badge>
                                                    </HStack>
                                                </Box>
                                                <AccordionIcon display={{ base: "none", md: "block" }} />
                                            </AccordionButton>
                                            <AccordionPanel pb={4}>
                                                <VStack alignItems="left">
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={600} color={'green.400'} w={{ base: 'fit-content', lg: "fit-content" }}>Confirmed: Travel miles is added to your account.</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction ID: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>1234-5678-9012-3456</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction date: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>26-08-21</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Amount spent: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>USD 317.79</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Remarks: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>30% Cashback For Every S$1 Spent</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Card: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>7063</Text>
                                                    </HStack>

                                                </VStack>
                                            </AccordionPanel>
                                        </AccordionItem>
                                    )
                                })}
                            </Accordion>
                        </TabPanel>
                        {cards.map((card) => {
                            const data = filteredTransactions.filter((el) => {
                                return el.card_pan === card;
                            })
                            return (
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                                    animate={{ opacity: 1 }}
                                    exit={{ opacity: 0 }} key={card}>
                                    <Accordion allowMultiple w="full">
                                        {data.map((transaction) => {
                                            return (
                                                <AccordionItem key={transaction.id}>
                                                    <AccordionButton p={{ base: 1, lg: 2 }}>
                                                        <Box as="span" flex='1' textAlign='left'>
                                                            <HStack pr={{ base: 0, lg: 8 }}>
                                                                <Text display={{ base: "none", md: "block" }}>{(transaction.mcc == '5499' | transaction.mcc == '5411') ? <CiShoppingBasket size={20} color='gray' /> : (transaction.mcc == '5541' | transaction.mcc == '5542') ? <CiDeliveryTruck size={20} color='gray' /> : (transaction.mcc == '5499') ? <CiForkAndKnife size={20} color='gray' /> : (transaction.mcc == '4121' | transaction.mcc == '5734' | transaction.mcc == '6540' | transaction.mcc == '4111') ? <CiRoute size={20} color='gray' /> : (transaction.mcc == '5999' | transaction.mcc == '5964' | transaction.mcc == '5691' | transaction.mcc == '5311' | transaction.mcc == '5411' | transaction.mcc == '5399' | transaction.mcc == '5311') ? <CiShoppingBasket size={20} color='gray' /> : <CiWallet size={20} color='gray' />}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 14, lg: 20 }}>{transaction.transaction_date}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 100, lg: 200 }}>{transaction.merchant}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 110, lg: 200 }}>Shopee 11.11 Sale</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 100, lg: 100 }}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                                <Spacer />
                                                                <Badge fontSize='xs' colorScheme="green" py={1} px={2} >+100 Miles</Badge>
                                                            </HStack>
                                                        </Box>
                                                        <AccordionIcon display={{ base: "none", md: "block" }} />
                                                    </AccordionButton>
                                                    <AccordionPanel pb={4}>
                                                        <VStack alignItems="left">
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={600} color={'green.400'} w={{ base: 'fit-content', lg: "fit-content" }}>Confirmed: Travel miles is added to your account.</Text>
                                                            </HStack>
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction ID: </Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>1234-5678-9012-3456</Text>
                                                            </HStack>
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction date: </Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>26-08-21</Text>
                                                            </HStack>
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Amount spent: </Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>USD 317.79</Text>
                                                            </HStack>
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Remarks: </Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>30% Cashback For Every S$1 Spent</Text>
                                                            </HStack>
                                                            <HStack>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Card: </Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>7063</Text>
                                                            </HStack>

                                                        </VStack>
                                                    </AccordionPanel>
                                                </AccordionItem>
                                            )
                                        })}
                                    </Accordion>
                                </TabPanel>
                            )
                        })}
                    </TabPanels>
                </Tabs>
                <Divider />
            </VStack>
        </Navbar>
    );
}


