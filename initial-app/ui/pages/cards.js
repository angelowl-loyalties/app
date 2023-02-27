import {
    Box,
    Divider,
    Heading,
    Select,
    HStack,
    Spacer,
    Switch,
    Tab,
    Table,
    TabList,
    TabPanel,
    TabPanels,
    Tabs,
    Tbody,
    Td,
    Text,
    Th,
    Thead,
    Tr,
    VStack,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { FaCcAmex, FaCcDiscover, FaCcMastercard, FaCcVisa, FaCreditCard } from 'react-icons/fa';
import { GiLibertyWing, GiShoppingBag } from 'react-icons/gi';
import { IoDiamond } from 'react-icons/io5';
import { MdOutlineFlightTakeoff } from 'react-icons/md';

import Navbar from '../components/Navbar';

export default function Cards() {
    const router = useRouter()
    const [toggle, setToggle] = useState(true)


    const style = {
        height: 300,
    };

    const interactivity = {
        mode: "cursor",
        actions: [
            {
                position: { x: [0, 1], y: [0, 1] },
                type: "seek",
                frames: [0, 120]
            }
        ],
    };


    const cards = [{
        'id': '4a8aa316-01fb-448f-bf37-f0a4ec4774a6',
        'card_id': '5c3dd517-adec-424d-ae8c-aa9c4948d496',
        'card_pan': '6771-8930-6970-2407',
        'card_type': 'scis_freedom',
        'created_at': '23-08-2021  9:09:40 AM',
        'updated_at': '23-08-2021  9:09:40 AM'
    },
    {
        'id': '4a8aa316-01fb-448f-bf37-f0a4ec4774a6',
        'card_id': '5d2123f3-9d5b-499b-90db-c0aec95432c2',
        'card_pan': '4775-8833-1918-5512',
        'card_type': 'scis_platinummiles',
        'created_at': '23-08-2021  7:28:14 AM',
        'updated_at': '23-08-2021  7:28:14 AM'
    },
    {
        'id': '4a8aa316-01fb-448f-bf37-f0a4ec4774a6',
        'card_id': '899fdb50-3a73-4cb3-8d75-134bc24fc3e9',
        'card_pan': '5380-3907-2820-7063',
        'card_type': 'scis_premiummiles',
        'created_at': '23-08-2021  8:46:10 AM',
        'updated_at': '23-08-2021  8:46:10 AM'
    },
    {
        'id': '4a8aa316-01fb-448f-bf37-f0a4ec4774a6',
        'card_id': '35b05f29-2129-4894-a786-f60b3e98e722',
        'card_pan': '4193-3687-9689-7635',
        'card_type': 'scis_shopping',
        'created_at': '23-08-2021  7:58:34 AM',
        'updated_at': '23-08-2021  7:58:34 AM'
    }
    ]

    var amex, visa, mastercard, discover, others, shopping, premium, platinum, freedom;


    amex = visa = mastercard = discover = others = shopping = premium = platinum = freedom = [];
    amex = cards.filter((el) => {
        return el.card_pan.charAt(0) == 3;
    })

    visa = cards.filter((el) => {
        return el.card_pan.charAt(0) == 4;
    })

    mastercard = cards.filter((el) => {
        return el.card_pan.charAt(0) == 5;
    })

    discover = cards.filter((el) => {
        return el.card_pan.charAt(0) == 6;
    })

    others = cards.filter((el) => {
        return el.card_pan.charAt(0) == 7;
    })

    shopping = cards.filter((el) => {
        return el.card_type == 'scis_shopping';
    })

    premium = cards.filter((el) => {
        return el.card_type == 'scis_premiummiles';
    })

    platinum = cards.filter((el) => {
        return el.card_type == 'scis_platinummiles';
    })

    freedom = cards.filter((el) => {
        return el.card_type == 'scis_freedom';
    })

    const [data, setData] = useState([cards, amex, visa, mastercard, discover, others])

    function handleToggle() {
        if (toggle) {
            setData([cards, shopping, premium, platinum, freedom])
        } else {
            setData([cards, amex, visa, mastercard, discover, others])
        }
        setToggle(!toggle)
    }

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
                return <FaCcAmex size={23} />
            case '4':
                return <FaCcVisa size={23} />
            case '5':
                return <FaCcMastercard size={23} />
            case '6':
                return <FaCcDiscover size={23} />
            default:
                return <FaCreditCard size={23} />
        }
    }

    return (
        <Navbar>
            <VStack alignItems='start' w="full">
                <HStack mb={{base: 4, lg: 6}}>
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
                            display={{base: "inline-block", lg: "none"}}
                            placeholder='All'>
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
                        <Box p={2} bgColor="gray.100" borderRadius="xl" display={{base: "none", lg: "inline-block"}}>
                            <HStack>
                                <TabList maxW={{ base: "xs", md: "2xl" }} >
                                    <Tab borderRadius='lg'>
                                        <Text mx={1} textStyle="tab">
                                            All
                                        </Text>
                                    </Tab>
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
                        <Text fontSize='sm' fontWeight={600} color={'gray.500'} lineHeight='7'>Card issuer</Text>
                        <Switch defaultChecked onChange={handleToggle} colorScheme='purple' size='md' />
                    </HStack>

                    <TabPanels >
                        {data && data.map((cards1) => {
                            return (
                                <TabPanel p={{base: 0, lg: 4}} mt={{base: 4, lg: 0}} key={cards1[0] && cards1[0].card_id}>
                                    <Table size='sm'>
                                        <Thead>
                                            <Tr>
                                                <Th></Th>
                                                <Th>PAN</Th>
                                                <Th>Card ID</Th>
                                                <Th>Type</Th>
                                                <Th>Created</Th>
                                                <Th>Updated</Th>
                                            </Tr>
                                        </Thead>
                                        <Tbody>
                                            {cards1 && cards1.map((card) => {
                                                return (
                                                    <>
                                                        <Tr key={card && card.card_id}>
                                                            <Td><Text fontSize='xs' color='gray.500'>{card && renderIssuer(card.card_pan)}</Text></Td>
                                                            <Td><Text fontSize='xs'>{card && card.card_pan.substring(card.card_pan.length - 4)}</Text></Td>
                                                            <Td><Text fontSize='xs' noOfLines={1} overflow="hidden">{card && card.card_id}</Text></Td>
                                                            <Td><Text fontSize='xs' color='gray.500'>{card && renderProgram(card.card_type)}</Text></Td>
                                                            <Td><Text fontSize='xs'>{card && card.created_at.split(" ")[0].replace('202', '2')}</Text></Td>
                                                            <Td><Text fontSize='xs'>{card && card.updated_at.split(" ")[0].replace('202', '2')}</Text></Td>
                                                        </Tr>
                                                    </>
                                                )
                                            })}
                                        </Tbody>
                                    </Table>
                                </TabPanel>
                            );
                        })
                        }
                    </TabPanels>
                </Tabs>
                <Divider />
            </VStack>
        </Navbar>
    );
}

