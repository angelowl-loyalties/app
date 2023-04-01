import { Search2Icon } from "@chakra-ui/icons";
import {
    Box,
    Card,
    CardBody,
    Divider,
    Heading,
    HStack,
    Select,
    Image,
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
    Stack,
    CardFooter,
    ButtonGroup,
    Button,
    Grid,
} from "@chakra-ui/react";
import { GiLibertyWing, GiShoppingBag } from "react-icons/gi";
import { IoDiamond } from "react-icons/io5";
import { MdOutlineFlightTakeoff } from "react-icons/md";
import { useSession } from "next-auth/react";
import { useRouter } from "next/router";
import Loading from "./loading";

import { useEffect, useState } from "react";
import Navbar from "../components/Navbar";

import axios from 'axios';


export default function Campaigns() {
    const router = useRouter();

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
            console.log(status);
            return;
        }
        axios
            .get(`https://itsag1t2.com/campaign`, {
                headers: { Authorization: session.id },
            })
            .then((response) => {
                console.log(response.data.data)
                setCampaigns(response.data.data);
            });
        setLoading(false);
    }, [session]);

    return (
        <>
            {loading ? (
                <Loading />
            ) : (
                <Navbar user>
                    <VStack alignItems="start" w="full">
                        <HStack mb={{ base: 4, lg: 6 }}>
                            <VStack alignItems="start">
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
                                        {campaigns.filter((campaign) => campaign.reward_program == "Shopping").map((campaign) => {
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
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl">
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
                                        {campaigns.filter((campaign) => campaign.reward_program == "PremiumMiles").map((campaign) => {
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
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl">
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
                                        {campaigns.filter((campaign) => campaign.reward_program == "PlatinumMiles").map((campaign) => {
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
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl">
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
                                        {campaigns.filter((campaign) => campaign.reward_program == "Freedom").map((campaign) => {
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
                                                <Card maxW={{ base: '100%', lg: '100%' }} borderRadius="xl">
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
                    </VStack>
                </Navbar>
            )}{" "}
        </>
    );
}
