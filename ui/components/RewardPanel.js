import {
    Accordion,
    AccordionButton,
    AccordionIcon,
    AccordionItem,
    AccordionPanel,
    Badge,
    Box,
    Button,
    Center,
    Container,
    HStack,
    Spacer,
    Spinner,
    TabPanel,
    Text,
    VStack,
} from '@chakra-ui/react';
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { CiDeliveryTruck, CiForkAndKnife, CiRoute, CiShoppingBasket, CiWallet } from 'react-icons/ci';

function RewardPanel(props) {
    const router = useRouter()
    const [cards, setCards] = useState({})
    const [filteredTransactions, setFilteredTransactions] = useState([])
    const [loading, setLoading] = useState(true)
    var runningTotal = 0
    const [pageNum, setPageNum] = useState(1)
    const [total, setTotal] = useState(0)
    const [data, setData] = useState([])
    const [cardType, setCardType] = useState([])

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        }
    });

    const paginate = () => {
        console.log(pageNum)
        try {
            axios.get(`https://itsag1t2.com/reward/${props.card.id}?page_no=2`, { headers: { Authorization: session.id } })
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
                    setData(prevTransactions => [...prevTransactions, ...truncatedTransactions])
                }).catch((error) => {
                    console.log(error)
                })
        } catch (error) {
            setPageNum(-1)
            console.log(error)
        }
    }

    useEffect(() => {
        axios.get(`https://itsag1t2.com/reward/${props.card.id}?page_no=${pageNum}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setPageNum(pageNum + 1)
                if (response.data.data.length < 20) {
                    setPageNum(-1)
                }
                const truncatedTransactions = response.data.data.map(transaction => ({
                    ...transaction,
                    amount: parseFloat(transaction.amount.toFixed(2))
                }));
                setData(prevTransactions => [...prevTransactions, ...truncatedTransactions])
            }).catch((error) => {
                console.log(error)
            })

        axios.get(`https://itsag1t2.com/reward/total/${props.card.id}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setTotal(response.data.data)
            }).catch((error) => {
                console.log(error)
            })

        axios.get(`https://itsag1t2.com/card/type?${props.card.card_type}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setCardType(response.data.data[0])
                setLoading(false)
            }).catch((error) => {
                console.log(error)
                setLoading(false)
            })
    }, [])
    runningTotal = total

    return (
        <>
            {loading ? <></> :
                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }} key={props.card.id}>
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
                            {data.sort((a, b) => new Date(b.transaction_date) - new Date(a.transaction_date)).map((transaction) => {
                                const temp = runningTotal
                                runningTotal = runningTotal - transaction.reward_amount
                                return (
                                    <AccordionItem key={transaction.id}>
                                        <AccordionButton p={{ base: 1, lg: 2 }}>
                                            <Box as="span" flex='1' textAlign='left'>
                                                <HStack pr={{ base: 0, lg: 8 }}>
                                                    <Text display={{ base: "none", md: "block" }}>{(transaction.mcc == '5499' | transaction.mcc == '5411') ? <CiShoppingBasket size={20} color='gray' /> : (transaction.mcc == '5541' | transaction.mcc == '5542') ? <CiDeliveryTruck size={20} color='gray' /> : (transaction.mcc == '5499') ? <CiForkAndKnife size={20} color='gray' /> : (transaction.mcc == '4121' | transaction.mcc == '5734' | transaction.mcc == '6540' | transaction.mcc == '4111') ? <CiRoute size={20} color='gray' /> : (transaction.mcc == '5999' | transaction.mcc == '5964' | transaction.mcc == '5691' | transaction.mcc == '5311' | transaction.mcc == '5411' | transaction.mcc == '5399' | transaction.mcc == '5311') ? <CiShoppingBasket size={20} color='gray' /> : <CiWallet size={20} color='gray' />}</Text>
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 14, lg: 40 }}>{transaction.transaction_date}</Text>
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 100, lg: 220 }}>{transaction.merchant}</Text>
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 120, lg: 180 }}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{ base: 120, lg: 130 }}>{Math.floor(temp)}</Text>
                                                    <Spacer />
                                                    <Badge fontSize='xs' colorScheme={transaction.reward_amount > 0 ? "green" : "gray"} py={1} px={2}>{transaction.reward_amount > 0 ? `+${Math.floor(transaction.reward_amount)} ${cardType.reward_unit}` : `0 ${cardType.reward_unit}`}</Badge>
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
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>Reward Programme: </Text>
                                                    <Text fontSize='xs' fontWeight={500} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>{`${transaction.card_type.replace("_", " ").toUpperCase()}`}</Text>
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
                            <Center display={pageNum == -1 ? "none" : "block"} textAlign="center" mt={5}><Button fontSize="xs" onClick={paginate}>View more</Button></Center>
                        </>
                    }
                </TabPanel>}</>

    );
}

export default RewardPanel;