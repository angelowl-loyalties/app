import { Search2Icon } from '@chakra-ui/icons';
import {
    Box,
    Card,
    CardBody,
    Divider,
    Heading,
    HStack,
    Input,
    InputGroup,
    InputLeftElement,
    ListItem,
    Spacer,
    Tab,
    TabList,
    TabPanel,
    TabPanels,
    Tabs,
    Text,
    UnorderedList,
    VStack,
} from '@chakra-ui/react';
import Image from 'next/image';
import { useEffect } from 'react';
import { GiLibertyWing, GiShoppingBag } from 'react-icons/gi';
import { IoDiamond } from 'react-icons/io5';
import { MdOutlineFlightTakeoff } from 'react-icons/md';

import Navbar from '../components/Navbar';

export default function Campaigns() {
    return (
        <Navbar>
            <VStack alignItems='start' w="full">
                <HStack mb={8}>
                    <VStack alignItems='start'>
                        <Heading fontWeight='bold' fontSize='2xl' pb={0}>Payment campaigns</Heading>
                        <Text fontSize='sm' fontWeight={600} color={'gray.500'} lineHeight='4'>
                            Supercharge your credit cards and get rewarded when you spend
                        </Text>
                    </VStack>
                </HStack>
                <Tabs variant='solid-rounded' colorScheme="purple" w="full">
                    <HStack>
                        <Box p={2} bgColor="gray.100" borderRadius="xl" >
                            <TabList>
                                <Tab fontSize='md' borderRadius='lg'>
                                    <GiShoppingBag size={23} />
                                    <Text ml={1} fontSize="sm">
                                        Shopping
                                    </Text>
                                </Tab>
                                <Tab fontSize='md' borderRadius='lg'>
                                    <MdOutlineFlightTakeoff size={23} />
                                    <Text ml={1} fontSize="sm">
                                        PremiumMiles
                                    </Text>
                                </Tab>
                                <Tab fontSize='md' borderRadius='lg'>
                                    <IoDiamond size={23} />
                                    <Text ml={1} fontSize="sm">
                                        PlatinumMiles
                                    </Text>
                                </Tab>
                                <Tab fontSize='md' borderRadius='lg'>
                                    <GiLibertyWing size={23} />
                                    <Text ml={1} fontSize="sm">
                                        Freedom
                                    </Text>
                                </Tab>
                            </TabList>
                        </Box>
                        <Spacer />
                        <InputGroup w="30%">
                            <InputLeftElement
                                pointerEvents='none'><Search2Icon color='gray.300' /></InputLeftElement>
                            <Input type='text' placeholder='Search' fontSize="sm" />
                        </InputGroup>
                    </HStack>

                    <TabPanels>
                        <TabPanel>
                            {[...Array(15).keys()].map((num) => {
                                return (
                                    <Card key={num} w="full" mb={4} border="1px" borderColor="gray.200">
                                        <HStack>
                                            <Box w={180} textAlign="-webkit-center">
                                                <Image src="/ascenda.png" height="150" width="150" style={{ objectFit: 'cover' }} alt='campaign image' />
                                            </Box>
                                            <CardBody px={0} py={4}>
                                                <VStack alignItems="start">
                                                    <Text fontSize='sm' fontWeight={600} color={'gray.900'}>
                                                        Supercharge your credit cards and get rewarded when you spend
                                                    </Text>
                                                    <UnorderedList px={6} fontSize="xs">
                                                        <ListItem>Lorem ipsum dolor sit amet</ListItem>
                                                        <ListItem>Consectetur adipiscing elit</ListItem>
                                                        <ListItem>Integer molestie lorem at massa</ListItem>
                                                        <ListItem>Facilisis in pretium nisl aliquet</ListItem>
                                                    </UnorderedList>
                                                </VStack>
                                            </CardBody>
                                        </HStack>
                                    </Card>
                                )
                            })}
                        </TabPanel>
                    </TabPanels>
                </Tabs>
                <Divider />
            </VStack>
        </Navbar>
    );
}


