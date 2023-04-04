import {
	AddIcon,
	CheckCircleIcon,
	DeleteIcon,
	EditIcon,
	MinusIcon,
} from "@chakra-ui/icons";
import {
	HStack,
	IconButton,
	Table,
	TableContainer,
	Tbody,
	Td,
	Text,
	Th,
	Thead,
	Tr,
	useDisclosure,
	useToast,
	VStack,
} from "@chakra-ui/react";
import axios from "axios";
import { useSession } from "next-auth/react";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import AddCampaigns from "../../components/AddCampaigns";
import AddExclusions from "../../components/AddExclusions";
import Loading from "../../components/Loading";
import Navbar from "../../components/Navbar";
import Upload from "../../components/UploadComponent";

export default function Home() {
	const {
		isOpen: isCampaignOpen,
		onOpen: onCampaignOpen,
		onClose: onCampaignClose,
	} = useDisclosure();
	const {
		isOpen: isExclusionOpen,
		onOpen: onExclusionOpen,
		onClose: onExclusionClose,
	} = useDisclosure();
	const toast = useToast();
	const router = useRouter();
	const [campaigns, setCampaigns] = useState([]);
	const [exclusions, setExclusions] = useState([]);
	const [loading, setLoading] = useState(true);
	const [refreshData, setRefreshData] = useState(false);

	const { data: session, status } = useSession({
		required: true,
		onUnauthenticated() {
			router.push("/admin/login");
		},
	});

	const refresh = () => {
		setRefreshData(!refreshData);
	};

	useEffect(() => {
		if (!session) {
			return;
		}

		axios
			.get(`https://itsag1t2.com/campaign`, {
				headers: { Authorization: session.id },
			})
			.then((response) => {
				setCampaigns(response.data.data);
				axios
					.get(`https://itsag1t2.com/exclusion`, {
						headers: { Authorization: session.id },
					})
					.then((response) => {
						setExclusions(response.data.data);
						setLoading(false);
					});
			})
			.finally((response) => {
				setLoading(false);
			});
	}, [session, refreshData]);

	return (
		<>
			{loading ? (
				<Loading />
			) : (
				<Navbar admin>
					<VStack alignItems="start" w="full">
						<HStack mb={4} w="full">
							<Text textStyle="head" mb={0}>
								Manage campaign(s)
							</Text>
							<IconButton
								onClick={onCampaignOpen}
								size="xs"
								variant="outline"
								colorScheme="green"
								icon={<AddIcon />}
							/>
						</HStack>
						<TableContainer w="full">
							<Table size="sm">
								<Thead>
									<Tr>
										<Th fontSize="x-small">Min</Th>
										<Th fontSize="x-small">From</Th>
										<Th fontSize="x-small">To</Th>
										<Th fontSize="x-small">R.program</Th>
										<Th fontSize="x-small">%/SGD</Th>
										<Th fontSize="x-small">MCC</Th>
										<Th fontSize="x-small">Merchant</Th>
										<Th fontSize="x-small">Base</Th>
										<Th fontSize="x-small">Foreign</Th>
										<Th fontSize="x-small">Actions</Th>
									</Tr>
								</Thead>
								<Tbody>
									{campaigns.map((campaign) => {
										return (
											<Tr key={campaign.id}>
												<Td fontSize="x-small">{campaign.min_spend}</Td>
												<Td fontSize="x-small">{campaign.start_date}</Td>
												<Td fontSize="x-small">{campaign.end_date}</Td>
												<Td fontSize="x-small">{campaign.reward_program}</Td>
												<Td fontSize="x-small">{campaign.reward_amount}</Td>
												<Td fontSize="x-small">{campaign.mcc}</Td>
												<Td fontSize="x-small">{campaign.merchant}</Td>
												<Td fontSize="x-small">
													{campaign.base_reward ? (
														<CheckCircleIcon color="green.400" />
													) : (
														<MinusIcon color="red.400" />
													)}
												</Td>
												<Td fontSize="x-small">
													{campaign.foreign_currency ? (
														<CheckCircleIcon color="green.400" />
													) : (
														<MinusIcon color="red.400" />
													)}
												</Td>
												<Td fontSize="x-small">
													<IconButton
														variant="ghost"
														size="xs"
														icon={<EditIcon color="blue" />}
													/>
													<IconButton
														size="xs"
														variant="ghost"
														icon={<DeleteIcon color="red" />}
													/>
												</Td>
											</Tr>
										);
									})}
								</Tbody>
							</Table>
						</TableContainer>
						<HStack pt={8} w="full">
							<Text textStyle="head" mb={0}>
								Manage exclusion(s)
							</Text>
							<IconButton
								onClick={onExclusionOpen}
								size="xs"
								variant="outline"
								colorScheme="green"
								icon={<AddIcon />}
							/>
						</HStack>
						<TableContainer w="full" pb={8}>
							<Table size="sm">
								<Thead>
									<Tr>
										<Th fontSize="x-small">MCC</Th>
										<Th fontSize="x-small">Valid from</Th>
									</Tr>
								</Thead>
								<Tbody>
									{exclusions.map((exclusion) => {
										return (
											<Tr key={exclusion.id}>
												<Td fontSize="x-small">{exclusion.mcc}</Td>
												<Td fontSize="x-small">{exclusion.valid_from}</Td>
												<Td fontSize="x-small">
													<IconButton
														variant="ghost"
														size="xs"
														icon={<EditIcon color="blue" />}
													/>
													<IconButton
														size="xs"
														variant="ghost"
														icon={<DeleteIcon color="red" />}
													/>
												</Td>
											</Tr>
										);
									})}
								</Tbody>
							</Table>
						</TableContainer>
						<Upload toast={toast} session={session} refresh={refresh} admin />
						<AddCampaigns
							isOpen={isCampaignOpen}
							onClose={onCampaignClose}
							toast={toast}
							session={session}
							refresh={refresh}
						/>
						<AddExclusions
							isOpen={isExclusionOpen}
							onClose={onExclusionClose}
							toast={toast}
							session={session}
							refresh={refresh}
						/>
					</VStack>
				</Navbar>
			)}
		</>
	);
}
