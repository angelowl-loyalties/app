import {
    Accordion,
    AccordionButton,
    AccordionItem,
    Box,
    Container,
    HStack,
    Select,
    Spacer,
    Tab,
    TabList,
    TabPanel,
    TabPanels,
    Tabs,
    Text,
    VStack,
} from '@chakra-ui/react';
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { FaCcAmex, FaCcDiscover, FaCcMastercard, FaCcVisa, FaCreditCard } from 'react-icons/fa';
import { GiLibertyWing, GiShoppingBag } from 'react-icons/gi';
import { IoDiamond } from 'react-icons/io5';
import { MdOutlineFlightTakeoff } from 'react-icons/md';

import Loading from '../components/Loading';
import Navbar from '../components/Navbar';

export default function Cards() {
    const router = useRouter()
    const [loading, setLoading] = useState(true)
    const [cards, setCards] = useState([])
    const [data, setData] = useState([])
    const [toggle, setToggle] = useState(true)
    var amex, visa, mastercard, discover, others, shopping, premium, platinum, freedom;
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
                setCards(response.data.data.CreditCards.map((card) => card.card_pan))

                amex = visa = mastercard = discover = others = shopping = premium = platinum = freedom = [];
                amex = response.data.data.CreditCards.filter((el) => {
                    return el.card_pan.charAt(0) == 3;
                })

                visa = response.data.data.CreditCards.filter((el) => {
                    return el.card_pan.charAt(0) == 4;
                })

                mastercard = response.data.data.CreditCards.filter((el) => {
                    return el.card_pan.charAt(0) == 5;
                })

                discover = response.data.data.CreditCards.filter((el) => {
                    return el.card_pan.charAt(0) == 6;
                })

                others = response.data.data.CreditCards.filter((el) => {
                    return el.card_pan.charAt(0) === "*";
                })

                shopping = response.data.data.CreditCards.filter((el) => {
                    return el.card_type == 'scis_shopping';
                })

                premium = response.data.data.CreditCards.filter((el) => {
                    return el.card_type == 'scis_premiummiles';
                })

                platinum = response.data.data.CreditCards.filter((el) => {
                    return el.card_type == 'scis_platinummiles';
                })

                freedom = response.data.data.CreditCards.filter((el) => {
                    return el.card_type == 'scis_freedom';
                })
                setData([amex, visa, mastercard, discover, others])
                setLoading(false)
            }).catch((error) => {
                console.log(error)
            })

    }, [session])

    function renderProgram(card_type) {
        if (card_type == 'scis_freedom') {
            return <GiLibertyWing size={20} />
        } else if (card_type == 'scis_platinummiles') {
            return <IoDiamond size={20} />
        } else if (card_type == 'scis_premiummiles') {
            return <MdOutlineFlightTakeoff size={20} />
        } else if (card_type == 'scis_shopping') {
            return <GiShoppingBag size={20} />
        }
    }

    function renderIssuer(card_pan) {
        switch (card_pan.charAt(0)) {
            case '3':
                return <FaCcAmex size={60} />
            case '4':
                return <FaCcVisa size={60} />
            case '5':
                return <FaCcMastercard size={60} />
            case '6':
                return <FaCcDiscover size={60} />
            default:
                return <FaCreditCard size={60} />
        }
    }

    return (
        <>
        { loading ? <Loading /> : 
        <Navbar user>
            <VStack alignItems='start' w="full">
                <HStack mb={{ base: 4, lg: 6 }}>
                    <VStack alignItems='start'>
                        <Text textStyle="title">View linked cards</Text>
                        <Text textStyle="subtitle" >
                            Supercharge your credit cards and get rewarded when you spend
                        </Text>
                    </VStack>
                </HStack>
                <Tabs variant='solid-rounded' colorScheme="purple" w="full">
                    <HStack>
                        <Select
                            w="25%"
                            fontSize="small"
                            display={{ base: "inline-block", lg: "none" }}>
                            <option>
                                {toggle ? 'AMEX' : 'Shopping'}
                            </option>
                            <option >
                                {toggle ? 'Visa' : 'PremiumMiles'}
                            </option>
                            <option>
                                {toggle ? 'Mastercard' : 'PlatinumMiles'}
                            </option>
                            <option>
                                {toggle ? 'Discover' : 'Freedom'}
                            </option>
                            <option>
                                Others
                            </option>
                        </Select>
                        <Box p={2} bgColor="gray.100" borderRadius="xl" display={{ base: "none", lg: "inline-block" }}>
                            <HStack>
                                <TabList maxW={{ base: "xs", md: "2xl" }} >
                                    <Tab borderRadius='lg'>
                                        {toggle ? <FaCcAmex size={23} /> : <GiShoppingBag size={23} />}
                                        <Text ml={1} textStyle="tab" isSelected>
                                            {toggle ? 'AMEX' : 'Shopping'}
                                        </Text>
                                    </Tab>
                                    <Tab borderRadius='lg'>
                                        {toggle ? <FaCcVisa size={23} /> : <MdOutlineFlightTakeoff size={23} />}
                                        <Text ml={1} textStyle="tab">
                                            {toggle ? 'Visa' : 'PremiumMiles'}
                                        </Text>
                                    </Tab>
                                    <Tab borderRadius='lg'>
                                        {toggle ? <FaCcMastercard size={23} /> : <IoDiamond size={23} />}
                                        <Text ml={1} textStyle="tab">
                                            {toggle ? 'Mastercard' : 'PlatinumMiles'}
                                        </Text>
                                    </Tab>
                                    <Tab borderRadius='lg'>
                                        {toggle ? <FaCcDiscover size={23} /> : <GiLibertyWing size={23} />}
                                        <Text ml={1} textStyle="tab">
                                            {toggle ? 'Discover' : 'Freedom'}
                                        </Text>
                                    </Tab>
                                    <Tab borderRadius='lg' display={!toggle ? 'none' : 'flex'}>
                                        <FaCreditCard size={23} />
                                        <Text ml={1} textStyle="tab">
                                            Others
                                        </Text>
                                    </Tab>
                                </TabList>
                            </HStack>
                        </Box>
                        <Spacer />
                    </HStack>

                    <TabPanels>
                        {data && data.map((cards1) => {
                            return (
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }} key={cards1 && cards1.card_id}>
                                    {cards1 && cards1.length == 0 ? <Container my={10} w="full" textAlign="-webkit-center">
                                        <svg width="64" height="41" viewBox="0 0 64 41" xmlns="http://www.w3.org/2000/svg">
                                            <g transform="translate(0 1)" fill="none" fillRule="evenodd">
                                                <ellipse fill="#f5f5f5" cx="32" cy="33" rx="32" ry="7"></ellipse>
                                                <g fillRule="nonzero" stroke="#d9d9d9">
                                                    <path d="M55 12.76L44.854 1.258C44.367.474 43.656 0 42.907 0H21.093c-.749 0-1.46.474-1.947 1.257L9 12.761V22h46v-9.24z"></path>
                                                    <path d="M41.613 15.931c0-1.605.994-2.93 2.227-2.931H55v18.137C55 33.26 53.68 35 52.05 35h-40.1C10.32 35 9 33.259 9 31.137V13h11.16c1.233 0 2.227 1.323 2.227 2.928v.022c0 1.605 1.005 2.901 2.237 2.901h14.752c1.232 0 2.237-1.308 2.237-2.913v-.007z" fill="#fafafa"></path>
                                                </g>
                                            </g>
                                        </svg>
                                        <Text fontSize='sm' fontWeight={500} color='#d9d9d9' lineHeight='7'>No cards found</Text></Container> :
                                        <Accordion allowMultiple w="full">
                                            {cards1 && cards1.map((card) => {
                                                return (
                                                    <AccordionItem key={card.id}>
                                                        <AccordionButton p={{ base: 1, lg: 3 }}>
                                                            <Text fontSize={{ base: 'xs', md: 'sm' }} color='gray.500' w={{ base: "fit-content", lg: "fit-content" }}>{card.card_pan && renderIssuer(card.card_pan)}</Text>
                                                            <Box textAlign="left" ml={{base: 5, md: 10}}>
                                                                <Text fontSize={{ base: 'xs', md: 'sm' }} noOfLines={1} overflow="hidden">{card && card.id}</Text>
                                                                <Text fontSize={{ base: 'xs', md: 'sm' }}>{card.card_pan && card.card_pan.substring(card.card_pan.length - 4)}</Text>
                                                                <Text fontSize={{ base: 'xs', md: 'sm' }} noOfLines={1} overflow="hidden">{card && card.card_type ? card.card_type : "Not eligible for rewards"} </Text>
                                                            </Box>
                                                        </AccordionButton>

                                                    </AccordionItem>
                                                )
                                            })}
                                        </Accordion>
                                    }
                                </TabPanel>
                            );
                        })
                        }
                    </TabPanels>
                </Tabs>
            </VStack>
        </Navbar>}</>
    );
}


