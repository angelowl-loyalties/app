import { QuestionIcon, Search2Icon } from '@chakra-ui/icons';
import {
    Box,
    Button,
    Card,
    CardBody,
    Divider,
    Grid,
    Heading,
    HStack,
    IconButton,
    Image,
    Input,
    InputGroup,
    InputLeftElement,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    Select,
    Spacer,
    Tab,
    Table,
    TableContainer,
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
    useDisclosure,
    VStack,
} from '@chakra-ui/react';
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { GiLibertyWing, GiShoppingBag } from 'react-icons/gi';
import { IoDiamond } from 'react-icons/io5';
import { MdOutlineFlightTakeoff } from 'react-icons/md';

import Loading from '../components/Loading';
import Navbar from '../components/Navbar';


export default function Campaigns() {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const router = useRouter();
    const [exclusions, setExclusions] = useState([]);
    const [loading, setLoading] = useState(true);
    const [campaigns, setCampaigns] = useState([]);
    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login");
        },
    });
    useEffect(() => {
        if (!session) {
            return;
        }
        axios
            .get(`https://itsag1t2.com/campaign`, {
                headers: { Authorization: session.id },
            }).then((response) => {
                console.log(response.data.data)
                setCampaigns(response.data.data);
                axios.get(`https://itsag1t2.com/exclusion`, {
                    headers: { Authorization: session.id },
                }).then((response) => {
                    console.log(response.data.data)
                    setExclusions(response.data.data);
                    setLoading(false);
                });
            });


    }, [session]);

    return (
        <>
            {loading ? (
                <Loading />
            ) : (
                <Navbar user>
                    <VStack alignItems="start" w="full">
                        <HStack mb={{ base: 4, lg: 6 }} >
                            <VStack alignItems="start" >
                                <Text textStyle="title">Payment campaigns</Text>
                                <Text textStyle="subtitle">
                                    Supercharge your credit cards and get rewarded when you spend
                                </Text>
                            </VStack>
                        </HStack>
                        <Tabs variant="solid-rounded" colorScheme="purple" w="full">
                            <HStack>
                                <Select
                                    w="25%"
                                    fontSize="small"
                                    display={{ base: "inline-block", lg: "none" }}
                                    placeholder="All"
                                >
                                    <option>Shopping</option>
                                    <option>PremiumMiles</option>
                                    <option>PlatinumMiles</option>
                                    <option>Freedom</option>
                                </Select>
                                <Box
                                    p={2}
                                    bgColor="gray.100"
                                    borderRadius="xl"
                                    display={{ base: "none", lg: "inline-block" }}
                                >
                                    <TabList>
                                        <Tab fontSize="md" borderRadius="lg">
                                            <GiShoppingBag size={23} />
                                            <Text ml={1} textStyle="tab">
                                                Shopping
                                            </Text>
                                        </Tab>
                                        <Tab fontSize="md" borderRadius="lg">
                                            <MdOutlineFlightTakeoff size={23} />
                                            <Text ml={1} textStyle="tab">
                                                PremiumMiles
                                            </Text>
                                        </Tab>
                                        <Tab fontSize="md" borderRadius="lg">
                                            <IoDiamond size={23} />
                                            <Text ml={1} textStyle="tab">
                                                PlatinumMiles
                                            </Text>
                                        </Tab>
                                        <Tab fontSize="md" borderRadius="lg">
                                            <GiLibertyWing size={23} />
                                            <Text ml={1} textStyle="tab">
                                                Freedom
                                            </Text>
                                        </Tab>
                                    </TabList>
                                </Box>
                                <IconButton size="lg" aria-label='Excluded' variant="ghost" onClick={onOpen} icon={<QuestionIcon color="gray.500" />} />
                                <Spacer />
                                <InputGroup w="30%">
                                    <InputLeftElement pointerEvents="none">
                                        <Search2Icon color="gray.300" />
                                    </InputLeftElement>
                                    <Input type="text" placeholder="Search" fontSize="sm" />
                                </InputGroup>
                            </HStack>

                            <TabPanels>
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }}>
                                    <Grid templateColumns={{ base: 'repeat(1, 1fr)', md: 'repeat(3, 1fr)' }} gap={5}>
                                        {campaigns.filter((campaign) => campaign.reward_program == "scis_shopping").map((campaign) => {
                                            const start = new Date(campaign.start_date);
                                            const end = new Date(campaign.end_date);

                                            const start_date = start.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });

                                            const end_date = end.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });
                                            return (
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl" key={campaigns.id}>
                                                    <CardBody>
                                                        <Image
                                                            src={`https://picsum.photos/seed/${Math.random()}/400/250`}
                                                            display={{ base: "none", md: "block" }}
                                                            w={{ base: "100%", lg: "100%" }}
                                                            alt='Green double couch with wooden legs'
                                                            borderRadius='lg'
                                                        />
                                                        <Heading size='md' py={4}>{campaign.name}</Heading>
                                                        <Text textAlign="justify" fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                            <Text fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                                Spend with your <b>SCIS {campaign.reward_program} Card</b> and earn rewards! <u>From {start_date} to {end_date}</u>,
                                                                earn <b>{campaign.reward_amount} {campaign.reward_program == "Freedom" ? "% Cashback" : campaign.reward_program == "PlatinumMiles" || campaign.reward_program == "PremiumMiles"
                                                                    ? "Miles/SGD" : campaign.reward_program == "Shopping" ? "Point(s) / SGD" : ""}</b> rewards when you make purchases at {campaign.merchant}.
                                                            </Text>
                                                        </Text>
                                                        <Text fontSize={12} color="blue.600" py={3} fontWeight={600} w={{ base: 'fit-content', lg: "fit-content" }}>Available for a limited time{campaign.foreign_currency ? "for foreign transactions only. T&C applies." : "."}</Text>
                                                    </CardBody>
                                                </Card>
                                            );
                                        })}
                                    </Grid>
                                </TabPanel>
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }}>
                                    <Grid templateColumns={{ base: 'repeat(1, 1fr)', md: 'repeat(3, 1fr)' }} gap={5}>
                                        {campaigns.filter((campaign) => campaign.reward_program == "scis_premiummiles").map((campaign) => {
                                            const start = new Date(campaign.start_date);
                                            const end = new Date(campaign.end_date);

                                            const start_date = start.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });

                                            const end_date = end.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });
                                            return (
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl" key={campaigns.id}>
                                                    <CardBody>
                                                        <Image
                                                            src={`https://picsum.photos/seed/${Math.random()}/400/250`}
                                                            display={{ base: "none", md: "block" }}
                                                            w={{ base: "100%", lg: "100%" }}
                                                            alt='Green double couch with wooden legs'
                                                            borderRadius='lg'
                                                        />
                                                        <Heading size='md' py={4}>{campaign.name}</Heading>
                                                        <Text textAlign="justify" fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                            <Text fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                                Spend with your <b>SCIS {campaign.reward_program} Card</b> and earn rewards! <u>From {start_date} to {end_date}</u>,
                                                                earn <b>{campaign.reward_amount} {campaign.reward_program == "Freedom" ? "% Cashback" : campaign.reward_program == "PlatinumMiles" || campaign.reward_program == "PremiumMiles"
                                                                    ? "Miles/SGD" : campaign.reward_program == "Shopping" ? "Point(s) / SGD" : ""}</b> rewards when you make purchases at {campaign.merchant}.
                                                            </Text>
                                                        </Text>
                                                        <Text fontSize={12} color="blue.600" py={3} fontWeight={600} w={{ base: 'fit-content', lg: "fit-content" }}>Available for a limited time{campaign.foreign_currency ? "for foreign transactions only. T&C applies." : "."}</Text>
                                                    </CardBody>
                                                </Card>
                                            );
                                        })}
                                    </Grid>
                                </TabPanel>
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }}>
                                    <Grid templateColumns={{ base: 'repeat(1, 1fr)', md: 'repeat(3, 1fr)' }} gap={5}>
                                        {campaigns.filter((campaign) => campaign.reward_program == "scis_platinummiles").map((campaign) => {
                                            const start = new Date(campaign.start_date);
                                            const end = new Date(campaign.end_date);

                                            const start_date = start.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });

                                            const end_date = end.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });
                                            return (
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl" key={campaigns.id}>
                                                    <CardBody>
                                                        <Image
                                                            src={`https://picsum.photos/seed/${Math.random()}/400/250`}
                                                            display={{ base: "none", md: "block" }}
                                                            w={{ base: "100%", lg: "100%" }}
                                                            alt='Green double couch with wooden legs'
                                                            borderRadius='lg'
                                                        />
                                                        <Heading size='md' py={4}>{campaign.name}</Heading>
                                                        <Text textAlign="justify" fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                            <Text fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                                Spend with your <b>SCIS {campaign.reward_program} Card</b> and earn rewards! <u>From {start_date} to {end_date}</u>,
                                                                earn <b>{campaign.reward_amount} {campaign.reward_program == "Freedom" ? "% Cashback" : campaign.reward_program == "PlatinumMiles" || campaign.reward_program == "PremiumMiles"
                                                                    ? "Miles/SGD" : campaign.reward_program == "Shopping" ? "Point(s) / SGD" : ""}</b> rewards when you make purchases at {campaign.merchant}.
                                                            </Text>
                                                        </Text>
                                                        <Text fontSize={12} color="blue.600" py={3} fontWeight={600} w={{ base: 'fit-content', lg: "fit-content" }}>Available for a limited time{campaign.foreign_currency ? " for foreign transactions only. T&C applies." : "."}</Text>
                                                    </CardBody>
                                                </Card>
                                            );
                                        })}
                                    </Grid>
                                </TabPanel>
                                <TabPanel p={{ base: 0, lg: 4 }} mt={{ base: 4, lg: 0 }}>
                                    <Grid templateColumns={{ base: 'repeat(1, 1fr)', md: 'repeat(3, 1fr)' }} gap={5}>
                                        {campaigns.filter((campaign) => campaign.reward_program == "scis_freedom").map((campaign) => {
                                            const start = new Date(campaign.start_date);
                                            const end = new Date(campaign.end_date);

                                            const start_date = start.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });

                                            const end_date = end.toLocaleString("en-SG", {
                                                day: "numeric",
                                                month: "short",
                                                year: "numeric",
                                                hour: "numeric",
                                                minute: "numeric",
                                                hour12: true,
                                            });
                                            return (
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl" key={campaigns.id}>
                                                    <CardBody>
                                                        <Image
                                                            src={`https://picsum.photos/seed/${Math.random()}/400/250`}
                                                            display={{ base: "none", md: "block" }}
                                                            w={{ base: "100%", lg: "100%" }}
                                                            alt='Green double couch with wooden legs'
                                                            borderRadius='lg'
                                                        />
                                                        <Heading size='md' py={4}>{campaign.name}</Heading>
                                                        <Text textAlign="justify" fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                            <Text fontSize={{ base: 'xs', md: 'sm' }} fontWeight={500} w={{ base: 'fit-content', lg: "fit-content" }}>
                                                                Spend with your <b>SCIS {campaign.reward_program} Card</b> and earn rewards! <u>From {start_date} to {end_date}</u>,
                                                                earn <b>{campaign.reward_amount} {campaign.reward_program == "Freedom" ? "% Cashback" : campaign.reward_program == "PlatinumMiles" || campaign.reward_program == "PremiumMiles"
                                                                    ? "Miles/SGD" : campaign.reward_program == "Shopping" ? "Point(s) / SGD" : ""}</b> rewards when you make purchases at {campaign.merchant}.
                                                            </Text>
                                                        </Text>
                                                        <Text fontSize={12} color="blue.600" py={3} fontWeight={600} w={{ base: 'fit-content', lg: "fit-content" }}>Available for a limited time{campaign.foreign_currency ? " for foreign transactions only. T&C applies." : "."}</Text>
                                                    </CardBody>
                                                </Card>
                                            );
                                        })}
                                    </Grid>
                                </TabPanel>
                            </TabPanels>
                        </Tabs>
                        <Divider />
                        <Modal onClose={onClose} size="lg" isOpen={isOpen}>
                            <ModalOverlay />
                            <ModalContent>
                                <ModalHeader>Excluded MCC transactions</ModalHeader>
                                <ModalCloseButton />
                                <ModalBody>
                                    <TableContainer w="fit-content">
                                        <Table size='sm'>
                                            <Thead>
                                                <Tr>
                                                    <Th fontSize="small">MCC</Th>
                                                    <Th fontSize="small">Valid from</Th>
                                                </Tr>
                                            </Thead>
                                            <Tbody>
                                                {exclusions.map((exclusion) => {
                                                    return (
                                                        <Tr key={exclusion.id}>
                                                            <Td fontSize="small">{exclusion.mcc}</Td>
                                                            <Td fontSize="small">{exclusion.valid_from}</Td>
                                                        </Tr>
                                                    );
                                                })}
                                            </Tbody>
                                        </Table>
                                    </TableContainer>
                                </ModalBody>
                                <ModalFooter>
                                    <Button variant='ghost' colorScheme="purple" mr={3} onClick={onClose}>
                                        Close
                                    </Button>
                                </ModalFooter>
                            </ModalContent>
                        </Modal>
                    </VStack>
                </Navbar>
            )}{" "}
        </>
    );
}
