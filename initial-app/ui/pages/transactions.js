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

export default function Transactions() {
    const router = useRouter()

    const cards = ['6771-8930-6970-2407', '4775-8833-1918-5512', '5380-3907-2820-7063']
    const [filteredTransactions, setFilteredTransactions] = useState([])


    useEffect(() => {
        const transactions = [
            {
                'id': '445579ab-a8b2-4e27-8ee2-01674703ca4d',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Brown  Robel and Rowe',
                'mcc': '1847',
                'currency': 'USD',
                'amount': 1.79,
                'sgd_amount': 0,
                'transaction_id': 'c6f8587f4e976e2a315340df48ab04b189d6bf4c46146b23cdfe0ba0c56c2730',
                'transaction_date': '25-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': '48f757e6-97f2-445e-9acf-e8f7198f9736',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Will  Herman and Bednar',
                'mcc': '8138',
                'currency': 'SGD',
                'amount': 32.19,
                'sgd_amount': 0,
                'transaction_id': 'e51e2e5dad7d829afd393ec292db858bac435acd9b87f84dc2caa23be7de04cf',
                'transaction_date': '24-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': 'a4d6b471-f264-4642-9f9d-68284d4b44fd',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Bashirian  Spencer and Braun',
                'mcc': '9602',
                'currency': 'SGD',
                'amount': 1.12,
                'sgd_amount': 0,
                'transaction_id': 'ceaff851c934d07294455cfd72b7c91b4c40342bcdaf1448402a73057485c305',
                'transaction_date': '23-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': 'd1a0c8ff-d1cb-4cec-83d1-0ad9b2699391',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Schuppe Inc',
                'mcc': '5046',
                'currency': 'USD',
                'amount': 69.13,
                'sgd_amount': 0,
                'transaction_id': '55f4b07c2c9f1b665ef5586ce1bc3e0bc8ab3df44f9f8f83e2d544aeb01a92d1',
                'transaction_date': '22-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': 'da8549d8-3bc4-4430-9f78-1fea3eecb1ec',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Rice  Hodkiewicz and Stamm',
                'mcc': '7948',
                'currency': 'SGD',
                'amount': 53.43,
                'sgd_amount': 0,
                'transaction_id': '6ef5692c4fd19db2b5a18015e88972abb53af8938695df252a22453175c89b8b',
                'transaction_date': '21-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': 'e5e12c8b-2f41-40e0-91da-25601f96a54d',
                'card_id': '037aed96-24bf-4b12-9227-7a2e35189243',
                'merchant': 'Gottlieb',
                'mcc': '2542',
                'currency': 'SGD',
                'amount': 3932.16,
                'sgd_amount': 0,
                'transaction_id': '68f5769c12f87ce8b9e7642d6bc1f51750a746f94db46220c514cbd1e4b7f967',
                'transaction_date': '21-08-21',
                'card_pan': '6771-8930-6970-2407',
                'card_type': 'scis_platinummiles'
            },
            {
                'id': 'fe3f1a94-304f-4acb-8d1a-91761519469e',
                'card_id': '0592c51c-7bdc-4c87-b797-975a1fbca7b8',
                'merchant': 'Grimes Schaefer and Ullrich',
                'mcc': '8877',
                'currency': 'USD',
                'amount': 304.13,
                'sgd_amount': 0,
                'transaction_id': 'f40534e3813c8de2373199a8a8b1bcd54592e09b8d7ccdb0eb5eceeb1ae0e02a',
                'transaction_date': '26-08-21',
                'card_pan': '4775-8833-1918-5512',
                'card_type': 'scis_premiummiles'
            },
            {
                'id': '2f479cbd-c20a-4e67-b12b-63fdef6c1202',
                'card_id': '3f0b5a3e-5e29-4c03-8f0d-73fb46ebf089',
                'merchant': 'Kerluke Group',
                'mcc': '2964',
                'currency': 'USD',
                'amount': 863.39,
                'sgd_amount': 0,
                'transaction_id': '2eb0b93080a97dd640ea81135e03940e41efcd68b3241f74f27d9ddfe6a2d856',
                'transaction_date': '29-08-21',
                'card_pan': '5380-3907-2820-7063',
                'card_type': 'scis_platinummiles'
            }
        ]

        setFilteredTransactions(transactions)
    }, [])

    const handleSearch = (e) => {
        const query = e.target.value

        setFilteredTransactions(transactions.filter())
    }

    return (
        <Navbar>
            <VStack alignItems='start' w="full">
                <HStack  mb={{base: 4, lg: 6}}>
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
                                const num = card.slice(-4)
                                switch (card.charAt(0)) {
                                    case '3':
                                        return <option fontSize='sm' borderRadius='lg'>(AMEX) {num}</option>
                                    case '4':
                                        return <option fontSize='sm' borderRadius='lg'>(VISA) {num}</option>
                                    case '5':
                                        return <option fontSize='sm' borderRadius='lg'>(MCC) {num}</option>
                                    case '6':
                                        return <option fontSize='sm' borderRadius='lg'>(DISC) {num}</option>
                                    default:
                                        return <option fontSize='sm' borderRadius='lg'>(OTHERS){num}</option>
                                }
                            })}
                        </Select>
                        <Box p={2} bgColor="gray.100" borderRadius="xl" display={{base: "none", lg: "inline-block"}}>
                            <TabList>
                                <Tab fontSize='sm' borderRadius='lg'><Text mx={1}>All</Text></Tab>
                                {cards.map((card) => {
                                    const num = card.slice(-4)
                                    switch (card.charAt(0)) {
                                        case '3':
                                            return <Tab fontSize='sm' borderRadius='lg'><HStack><FaCcAmex size={23} /><Text>{num}</Text></HStack></Tab>
                                        case '4':
                                            return <Tab fontSize='sm' borderRadius='lg'><HStack><FaCcVisa size={23} /><Text>{num}</Text></HStack></Tab>
                                        case '5':
                                            return <Tab fontSize='sm' borderRadius='lg'><HStack><FaCcMastercard size={23} /><Text>{num}</Text></HStack></Tab>
                                        case '6':
                                            return <Tab fontSize='sm' borderRadius='lg'><HStack><FaCcDiscover size={23} /><Text>{num}</Text></HStack></Tab>
                                        default:
                                            return <Tab fontSize='sm' borderRadius='lg'><HStack><FaCreditCard size={23} /><Text>{num}</Text></HStack></Tab>
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
                            const data = filteredTransactions.filter((el) => {
                                return el.card_pan === card;
                            })
                            return (
                                <TabPanel p={{base: 0, lg: 4}} mt={{base: 4, lg: 0}} initial={{ opacity: 0 }}
                                    animate={{ opacity: 1 }}
                                    exit={{ opacity: 0 }} key={card}>
                                    <Accordion allowMultiple w="full">
                                        {data.map((transaction) => {
                                            return (
                                                <AccordionItem key={transaction}>
                                                    <AccordionButton p={{base: 1, lg: 2}}>
                                                        <Box as="span" flex='1' textAlign='left'>
                                                            <HStack pr={{base: 0, lg: 8}}>
                                                                <Text display={{base: "none", md: "block"}}>{(transaction.mcc == '5499' | transaction.mcc == '5411') ? <CiShoppingBasket size={20} color='gray' /> : (transaction.mcc == '5541' | transaction.mcc == '5542') ? <CiDeliveryTruck size={20} color='gray' /> : (transaction.mcc == '5499') ? <CiForkAndKnife size={20} color='gray' /> : (transaction.mcc == '4121' | transaction.mcc == '5734' | transaction.mcc == '6540' | transaction.mcc == '4111') ? <CiRoute size={20} color='gray' /> : (transaction.mcc == '5999' | transaction.mcc == '5964' | transaction.mcc == '5691' | transaction.mcc == '5311' | transaction.mcc == '5411' | transaction.mcc == '5399' | transaction.mcc == '5311') ? <CiShoppingBasket size={20} color='gray' /> : <CiWallet size={20} color='gray' />}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{base: 14, lg: 20}}>{transaction.transaction_date}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{base: 100, lg: 200}}>{transaction.merchant}</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' noOfLines={1} overflow="hidden" w={{base: 110, lg: 200}}>Shopee 11.11 Sale</Text>
                                                                <Text fontSize='xs' fontWeight={500} color={'gray.500'} lineHeight='7' w={{base: 100, lg: 100}}>{"- " + transaction.currency + " " + transaction.amount}</Text>
                                                                <Spacer />
                                                                <Badge fontSize='xs' colorScheme="green" py={1} px={2} >+200 Miles</Badge>
                                                            </HStack>
                                                        </Box>
                                                        <AccordionIcon display={{base: "none", md: "block"}}/>
                                                    </AccordionButton>
                                                    <AccordionPanel pb={4}>

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


