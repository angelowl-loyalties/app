import { useToast } from "@chakra-ui/react";
import {
	HStack,
	Input,
	Text,
	VStack,
	FormControl,
	NumberInputField,
	NumberInput,
	NumberInputStepper,
	NumberDecrementStepper,
	NumberIncrementStepper,
	FormLabel,
	RadioGroup,
	Radio,
	Button,
} from "@chakra-ui/react";
import axios from "axios";
import { useSession } from "next-auth/react";

import { useState } from "react";

import Navbar from "../../components/Navbar";

export default function AddExclusions() {
	const [campaignName, setCampaignName] = useState("");
	const [minSpend, setMinSpend] = useState(0);
	const [startDate, setStartDate] = useState(new Date());
	const [endDate, setEndDate] = useState(new Date());
	const [rewardProgram, setRewardProgram] = useState("Shopping");
	const [rewardAmount, setRewardAmount] = useState(0);
	const [mcc, setMcc] = useState(0);
	const [foreignCurrency, setForeignCurrency] = useState(false);
	const [merchant, setMerchant] = useState("");

	const { data: session, status } = useSession({
		required: true,
		onUnauthenticated() {
			router.push("/login");
		},
	});

	const handleSubmit = (event) => {
		event.preventDefault();
		console.log(
			campaignName,
			minSpend,
			startDate,
			endDate,
			rewardProgram,
			rewardAmount,
			mcc,
			foreignCurrency,
			merchant
		);
	};
	const toast = useToast();

	const addExclusion = () => {
		toast({
			title: "In progress",
			description: "Please hold on while we create an exclusion",
			status: "info",
			duration: 9000,
			isClosable: true,
		});
		const body = {
			// name: campaignName,
			// min_spend: parseFloat(minSpend),
			valid_from: startDate + ":00Z",
			// end_date: endDate + ":00Z",
			// reward_program: rewardProgram,
			// reward_amount: parseInt(rewardAmount),
			mcc: parseInt(mcc),
			// merchant: merchant,
			// foreign_currency: foreignCurrency,
		};
		axios
			.post(`https://itsag1t2.com/exclusion`, body, {
				headers: {
					Authorization: session.id,
				},
			})
			.then((response) => {
				console.log(response);
				toast.closeAll();
				toast({
					title: "Success",
					description: `Exclusion created successfully`,
					status: "success",
					duration: 9000,
					isClosable: true,
				});
			})
			.catch((error) => {
				console.log(error);
				toast.closeAll();
				toast({
					title: "Error",
					description: "An error occurred while creating a exclusion",
					status: "error",
					duration: 9000,
					isClosable: true,
				});
			});
	};
	return (
		<Navbar admin>
			<VStack alignItems="start" w="full">
				<HStack mb={{ base: 4, lg: 6 }}>
					<VStack alignItems="start">
						<Text textStyle="title">Create Exclusions</Text>
						<Text textStyle="subtitle">Create Exclusions</Text>
					</VStack>
				</HStack>
				<HStack>
					<VStack>
						<form onSubmit={handleSubmit}>
							<FormControl as="fieldset">
								{/* <FormLabel mt={4}>Exclusion Name</FormLabel>

								<Input
									placeholder="Campaign Name"
									onChange={(event) => {
										setCampaignName(event.currentTarget.value);
									}}
								/> */}

								<FormControl isRequired>
									<FormLabel mt={4}>Start Date</FormLabel>
									<Input
										placeholder="Select Date and Time"
										size="md"
										type="datetime-local"
										onChange={(event) => {
											setStartDate(event.currentTarget.value);
										}}
									/>
								</FormControl>

								<FormLabel mt={4}>Applicable MCC</FormLabel>
								<NumberInput
									max={9999}
									min={1}
									defaultValue={0}
									onChange={(event) => {
										setMcc(event);
									}}
								>
									<NumberInputField />
									<NumberInputStepper>
										<NumberIncrementStepper />
										<NumberDecrementStepper />
									</NumberInputStepper>
								</NumberInput>

								<Button
									variantColor="teal"
									variant="outline"
									type="submit"
									width="full"
									mt={4}
									onClick={() => addExclusion()}
								>
									Add Exclusion
								</Button>
							</FormControl>
						</form>
					</VStack>
				</HStack>
			</VStack>
		</Navbar>
	);
}
