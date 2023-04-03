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
    Skeleton,
    Spacer,
    Spinner,
    TabPanel,
    Text,
    VStack,
} from '@chakra-ui/react';
import Fuse from 'fuse.js'
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { CiDeliveryTruck, CiForkAndKnife, CiRoute, CiShoppingBasket, CiWallet } from 'react-icons/ci';
import { CheckCircleIcon, CopyIcon } from '@chakra-ui/icons';

function RewardPanel(props) {
    const router = useRouter()
    const [cards, setCards] = useState({})
    const [filteredTransactions, setFilteredTransactions] = useState([])
    const [loading, setLoading] = useState(true)
    var runningTotal = 0
    const [pageNum, setPageNum] = useState(1)
    const [total, setTotal] = useState(0)
    const [data, setData] = useState()
    const [cardType, setCardType] = useState([])

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        }
    });

    const paginate = () => {
        try {
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
                    handleSearch(props.search)
                }).catch((error) => {
                    console.log(error)
                })
        } catch (error) {
            setPageNum(-1)
            console.log(error)
        }
    }

    useEffect(() => {
        if (!props.search) {
            setFilteredTransactions(data)
            return
        }
        handleSearch(props.search)
    }, [props.search])

    function handleSearch(search) {
        const options = {
            shouldSort: true,
            isCaseSensitive: false,
            keys: ["transaction_date", "id", "card_pan", "card_type", "currency",
                "merchant", "remarks", "sgd_amount", "transaction_id"]
        };
        const fuse = new Fuse(data, options);
        const pattern = search
        try {
            console.log(fuse.search(pattern))
            setFilteredTransactions(fuse.search(pattern).map((item) => item.item))
        } catch (error) {
            console.log("Skipped search in empty list")
        }
    }

    useEffect(() => {
        if (!props) {
            return
        }
        axios.get(`https://itsag1t2.com/reward/${props.card.id}?page_no=${pageNum}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setPageNum(pageNum + 1)
                const truncatedTransactions = response.data.data.map(transaction => ({
                    ...transaction,
                    amount: parseFloat(transaction.amount.toFixed(2))
                }));
                setData(prevTransactions => {
                    const transactionsArray = Array.isArray(prevTransactions) ? prevTransactions : [];
                    return [...transactionsArray, ...truncatedTransactions];
                });
                setFilteredTransactions(prevTransactions => {
                    const transactionsArray = Array.isArray(prevTransactions) ? prevTransactions : [];
                    return [...transactionsArray, ...truncatedTransactions];
                })

                if (response.data.data.length < 20) {
                    setPageNum(-1)
                    setLoading(false)
                } else {
                    setLoading(false)
                    paginate()
                }
            }).catch((error) => {
                console.log(error)
                setLoading(false)
            })

        axios.get(`https://itsag1t2.com/reward/total/${props.card.id}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setTotal(response.data.data)
            }).catch((error) => {
                console.log(error)
            })

        axios.get(`https://itsag1t2.com/card/type?${props.card.card_type}`, { headers: { Authorization: session.id } })
            .then((response) => {
                const card = response.data.data.find(card => card.card_type === props.card.card_type)
                setCardType(card)
            }).catch((error) => {
                console.log(error)
            })
    }, [props.card])

    runningTotal = total

    return (
        <>
            {loading ? <VStack>
                <Skeleton height='20px' />
                <Skeleton height='20px' />
                <Skeleton height='20px' />
            </VStack> :
                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }} key={props.card.id}>
                    {filteredTransactions == undefined ? <Container my={10} w="full" textAlign="-webkit-center">
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
                            {filteredTransactions && filteredTransactions.sort((a, b) => new Date(b.transaction_date) - new Date(a.transaction_date)).map((transaction) => {
                                const temp = runningTotal
                                runningTotal = runningTotal - transaction.reward_amount
                                return (
                                    <AccordionItem key={transaction.id}>
                                        <AccordionButton p={{ base: 1, lg: 2 }}>
                                            <Box as="span" flex='1' textAlign='left'>
                                                <HStack pr={{ base: 0, lg: 8 }}>
                                                    <Text display={{ base: "none", md: "block" }}><CiShoppingBasket size={20} color='gray' /></Text>
                                                    <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} color={'gray.700'} lineHeight='7' display={{ base: "block", lg: "none" }} w={{ base: 10, lg: 40 }}>{transaction.transaction_date.substring(5)}</Text>
                                                    <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} color={'gray.700'} lineHeight='7' display={{ base: "none", lg: "block" }} w={{ base: 14, lg: 40 }}>{transaction.transaction_date}</Text>
                                                    <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} color={'gray.700'} lineHeight='7' noOfLines={1} overflow="hidden" w={{ base: 62, lg: 220 }}>{transaction.merchant}</Text>
                                                    <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} color={'gray.700'} lineHeight='7' w={{ base: 70, lg: 180 }}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                    <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} color={'gray.700'} lineHeight='7' w={{ base: 62, lg: 130 }}>{Math.floor(temp) + " " + cardType.reward_unit}</Text>
                                                    <Spacer />
                                                    <Badge fontSize={{ base: 'xx-small', md: 'sm' }} colorScheme={transaction.reward_amount > 0 ? "green" : "gray"} py={1} px={2}>{transaction.reward_amount > 0 ? `+${Math.floor(transaction.reward_amount)} ${cardType.reward_unit}` : `0 ${cardType.reward_unit}`}</Badge>
                                                </HStack>
                                            </Box>
                                            <AccordionIcon display={{ base: "none", md: "block" }} />
                                        </AccordionButton>
                                        <AccordionPanel p={{base: 2, md: 5}} backgroundColor={transaction.reward_amount > 0 ? 'green.50' : "blue.50"}>
                                            <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={700} color={transaction.reward_amount > 0 ? 'green.400' : "blue.400"} w={{ base: 'fit-content', lg: "fit-content" }} alignItems="center" display="flex" >{transaction.reward_amount > 0 ? 'Confirmed' : 'Processed'}: {transaction.merchant} ({transaction.transaction_date}) <CheckCircleIcon ml={2} mb={1} color={'green.400'} display={transaction.reward_amount > 0 ? "block" : "none"} /></Text>
                                            <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>ID: {transaction.id}<Button size="xs" variant="unstyled" display={{base: "none", md: "block"}}><CopyIcon  /></Button></Text>
                                            <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>Reward Programme: {`${transaction.card_type.replace("_", " ").toUpperCase()}`}</Text>
                                            <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>Amount: {`${transaction.currency} ${transaction.amount}`}</Text>
                                            <Text fontSize={{ base: 'xx-small', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>Remarks: {`${transaction.remarks ? transaction.remarks : "Not applicable"}`}</Text>
                                            <Text fontSize={{ base: 7, md: 'sm' }} color={'gray.500'} w={{ base: 'fit-content', lg: "fit-content" }}>This transaction may be subject to additional fees or charges imposed by the merchant or card issuer.</Text>
                                        </AccordionPanel>
                                    </AccordionItem>
                                )
                            })}
                        </Accordion>
                            <Center display={pageNum >= -1 ? "none" : "block"} textAlign="center" mt={5}><Button fontSize="xs" onClick={paginate}>View more</Button></Center>
                        </>
                    }
                </TabPanel>}</>

    );
}

export default RewardPanel;