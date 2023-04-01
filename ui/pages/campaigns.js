import { Search2Icon } from "@chakra-ui/icons";
import {
	Box,
	Card,
	CardBody,
	Divider,
	Heading,
	HStack,
	Select,
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
} from "@chakra-ui/react";
import Image from "next/image";
import { GiLibertyWing, GiShoppingBag } from "react-icons/gi";
import { IoDiamond } from "react-icons/io5";
import { MdOutlineFlightTakeoff } from "react-icons/md";
import { useSession } from "next-auth/react";
import { useRouter } from "next/router";
import Loading from "./loading";

import { useEffect, useState } from "react";
import Navbar from "../components/Navbar";

export default function Campaigns() {
	const router = useRouter();
	const [loading, setLoading] = useState(true);
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
									{[...Array(15).keys()].map((num) => {
										return (
											<Card
												key={num}
												w="full"
												mb={4}
												border="1px"
												borderColor="gray.200"
											>
												<HStack>
													<Box w={180} textAlign="-webkit-center">
														<Image
															src="/ascenda.webp"
															height="150"
															width="150"
															style={{ objectFit: "cover" }}
															alt="campaign image"
														/>
													</Box>
													<CardBody px={0} py={4}>
														<VStack alignItems="start">
															<Text
																fontSize="sm"
																fontWeight={600}
																color={"gray.900"}
															>
																Supercharge your credit cards and get rewarded
																when you spend
															</Text>
															<UnorderedList px={6} fontSize="xs">
																<ListItem>Lorem ipsum dolor sit amet</ListItem>
																<ListItem>Consectetur adipiscing elit</ListItem>
																<ListItem>
																	Integer molestie lorem at massa
																</ListItem>
																<ListItem>
																	Facilisis in pretium nisl aliquet
																</ListItem>
															</UnorderedList>
														</VStack>
													</CardBody>
												</HStack>
											</Card>
										);
									})}
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
