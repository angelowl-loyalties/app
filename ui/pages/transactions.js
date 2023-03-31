import { Search2Icon } from '@chakra-ui/icons';
import { useRouter } from 'next/router';
import { useState, useEffect } from 'react';
import Loading from './loading';
import { useSession } from 'next-auth/react';
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
    Container,
    Button,
    Center,
} from '@chakra-ui/react';
import { CiDeliveryTruck, CiForkAndKnife, CiRoute, CiShoppingBasket, CiWallet } from 'react-icons/ci';
import { FaCcAmex, FaCcDiscover, FaCcMastercard, FaCcVisa, FaCreditCard } from 'react-icons/fa';

import Navbar from '../components/Navbar';
import axios from 'axios';

export default function Transactions() {
    const router = useRouter()
    const [cards, setCards] = useState({})
    const [filteredTransactions, setFilteredTransactions] = useState([])
    const [loading, setLoading] = useState(true)
    const [pageNum, setPageNum] = useState(1)
    const [total, setTotal] = useState(0)

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        }
    });

    useEffect(() => {
        if (!session){
            return
        }
        axios.get(`https://itsag1t2.com/user/${session.userId}`, { headers: { Authorization: session.id } })
            .then((response) => {
                console.log(response.data.data.CreditCards)
                setCards(response.data.data.CreditCards)

                response.data.data.CreditCards.map((card) => {
                    axios.get(`https://itsag1t2.com/reward/${card.id}?page_no=${pageNum}`, { headers: { Authorization: session.id } })
                        .then((response) => {
                            console.log(response.data.data)
                            setPageNum(pageNum + 1)
                            if (response.data.data.length < 20) {
                                setPageNum(-1)
                            }
                            const truncatedTransactions = response.data.data.map(transaction => ({
                                ...transaction,
                                amount: parseFloat(transaction.amount.toFixed(2))
                            }));
                            setFilteredTransactions(prevTransactions => [...prevTransactions, ...truncatedTransactions])
                        }).catch((error) => {
                            console.log(error)
                        })
                        
                        axios.get(`https://itsag1t2.com/reward/total/${card.id}`, { headers: { Authorization: session.id } })
                        .then((response) => {
                            setTotal(total + response.data.data)
                            setLoading(false)
                        }).catch((error) => {
                            console.log(error)
                            setLoading(false)
                        })
                })
            })


    }, [session])

    function getRunningTotal(reward) {
        const temp = total
        setTotal(total - reward)
    }

    const loadPage = (e) => {
        try{
            axios.get(`https://itsag1t2.com/reward/${e.target.id}?page_no=${pageNum}`, { headers: { Authorization: session.id } })
                .then((response) => {
                    console.log(response.data.data)
                    setPageNum(pageNum + 1)
                    if (response.data.data.length < 20) {
                        setPageNum(-1)
                    }
                    const truncatedTransactions = response.data.data.map(transaction => ({
                        ...transaction,
                        amount: parseFloat(transaction.amount.toFixed(2))
                    }));
                    setFilteredTransactions(prevTransactions => [...prevTransactions, ...truncatedTransactions])
                }).catch((error) => {
                    console.log(error)
                })
        } catch (e) {
            setPageNum(-1)
            console.log(e)
        }
    }

    const handleSearch = (e) => {
        // const query = e.target.value
        // setFilteredTransactions(transactions.filter())
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
                {/* TODO: Enable lazy loading */}
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
                                <Tab fontSize='sm' borderRadius='lg'><Text mx={1}>All</Text></Tab>
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
                        <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            exit={{ opacity: 0 }} key="all">
                            <Accordion allowMultiple w="full">
                                {filteredTransactions.sort((a, b) => new Date(b.transaction_date) - new Date(a.transaction_date)).map((transaction) => {
                                    return (
                                        <AccordionItem key={transaction.id}>
                                            <AccordionButton p={{ base: 1, lg: 2 }}>
                                                <Box as="span" flex='1' textAlign='left'>
                                                    <HStack pr={{ base: 0, lg: 8 }}>
                                                        <Text display={{ base: "none", md: "block" }}>{(transaction.mcc == '5499' | transaction.mcc == '5411') ? <CiShoppingBasket size={20} color='gray' /> : (transaction.mcc == '5541' | transaction.mcc == '5542') ? <CiDeliveryTruck size={20} color='gray' /> : (transaction.mcc == '5499') ? <CiForkAndKnife size={20} color='gray' /> : (transaction.mcc == '4121' | transaction.mcc == '5734' | transaction.mcc == '6540' | transaction.mcc == '4111') ? <CiRoute size={20} color='gray' /> : (transaction.mcc == '5999' | transaction.mcc == '5964' | transaction.mcc == '5691' | transaction.mcc == '5311' | transaction.mcc == '5411' | transaction.mcc == '5399' | transaction.mcc == '5311') ? <CiShoppingBasket size={20} color='gray' /> : <CiWallet size={20} color='gray' />}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 14, lg: 40 }}>{transaction.transaction_date}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 100, lg: 220 }}>{transaction.merchant}</Text>
                                                        {/* <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 110, lg: 180 }}>Shopee 11.11 Sale</Text> */}
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 120, lg: 180 }}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 120, lg: 130 }} onLoad= {() => setTotal(total - transaction.reward_amount)}>{total}</Text>
                                                        <Spacer />
                                                        <Badge fontSize='xs' colorScheme={transaction.reward_amount > 0 ? "green" : "gray"} py={1} px={2}>{transaction.reward_amount > 0 ? `+${transaction.reward_amount} points` : "0 points"}</Badge>
                                                    </HStack>
                                                </Box>
                                                <AccordionIcon display={{ base: "none", md: "block" }} />
                                            </AccordionButton>
                                            <AccordionPanel pb={4}>
                                                <VStack alignItems="left">
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={600} color={'green.400'} w={{ base: 'fit-content', lg: "fit-content" }}>Confirmed: {transaction.merchant}.</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction ID: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{transaction.id}</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Transaction date: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{transaction.transaction_date}</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Amount spent: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{`${transaction.currency} ${transaction.amount}`}</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Remarks: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{`${transaction.remarks ? transaction.remarks : "Not applicable"}`}</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Reard Programme: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{`${transaction.card_type.replace("_"," ").toUpperCase()}`}</Text>
                                                    </HStack>
                                                    <HStack>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Card: </Text>
                                                        <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{transaction.card_pan.slice(-4)}</Text>
                                                    </HStack>

                                                </VStack>
                                            </AccordionPanel>
                                        </AccordionItem>
                                    )
                                })}
                            </Accordion>
                            <Center display={pageNum == -1? "none" : "block"} textAlign="center" mt={5}><Button fontSize="xs" >Load more transactions</Button></Center>
                        </TabPanel>
                        {cards.map((card) => {
                            const data = filteredTransactions.filter((el) => {
                                return el.card_pan === card.card_pan;
                            })
                            return (
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                                    animate={{ opacity: 1 }}
                                    exit={{ opacity: 0 }} key={card.id}>
                                    {data.length == 0 ? <Container my={10} w="full" textAlign="-webkit-center">
                                        <svg width="64" height="41" viewBox="0 0 64 41" xmlns="http://www.w3.org/2000/svg">
                                            <g transform="translate(0 1)" fill="none" fillRule="evenodd">
                                                <ellipse fill="#f5f5f5" cx="32" cy="33" rx="32" ry="7"></ellipse>
                                                <g fillRule="nonzero" stroke="#d9d9d9">
                                                    <path d="M55 12.76L44.854 1.258C44.367.474 43.656 0 42.907 0H21.093c-.749 0-1.46.474-1.947 1.257L9 12.761V22h46v-9.24z"></path>
                                                    <path d="M41.613 15.931c0-1.605.994-2.93 2.227-2.931H55v18.137C55 33.26 53.68 35 52.05 35h-40.1C10.32 35 9 33.259 9 31.137V13h11.16c1.233 0 2.227 1.323 2.227 2.928v.022c0 1.605 1.005 2.901 2.237 2.901h14.752c1.232 0 2.237-1.308 2.237-2.913v-.007z" fill="#fafafa"></path>
                                                </g>
                                            </g>
                                        </svg>
                                        <Text fontSize='sm' fontWeight={500} color='#d9d9d9' lineHeight='7'>No transactions found</Text></Container> :
                                        <><Accordion allowMultiple w="full">
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
                                        <Center display={pageNum == -1? "none" : "block"} textAlign="center" mt={5}><Button fontSize="xs" >Load more</Button></Center>
                                        </>
                                        }
                                </TabPanel>
                            )
                        })}
                    </TabPanels>
                </Tabs>
            </VStack>
        </Navbar>}
        </>
    );
}


