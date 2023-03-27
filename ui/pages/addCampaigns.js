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
	FormControl,
	NumberInputField,
	NumberInput,
	NumberInputStepper,
	NumberDecrementStepper,
	NumberIncrementStepper,
	FormLabel,
	FormErrorMessage,
	FormHelperText,
	RadioGroup,
	Radio,
	Button,
} from "@chakra-ui/react";
import axios from "axios";

import { useState } from "react";

import Navbar from "../components/Navbar";

export default function AddCampaigns() {
	const [campaignName, setCampaignName] = useState("");
	const [minSpend, setMinSpend] = useState(0);
	const [startDate, setStartDate] = useState(new Date());
	const [endDate, setEndDate] = useState(new Date());
	const [rewardProgram, setRewardProgram] = useState("");
	const [rewardAmount, setRewardAmount] = useState(0);
	const [mcc, setMcc] = useState("");
	const [foreignCurrency, setForeignCurrency] = useState(false);
	const [merchant, setMerchant] = useState("");

	const jwtKey = ""
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
	const addCampaign = () => {
		console.log("Hello")

		const body = {
			name: campaignName,
			min_spend: minSpend,
			start: startDate,
			end: endDate,
			reward_program: rewardProgram,
			reward_amount: rewardAmount,
			mcc: mcc.toString(),
			merchant: merchant,
			foreign_currency: foreignCurrency,
		};
		console.log(body)
		axios
			.post(`https://itsag1t2.com/campaign`, body,{
				headers: {
					Authorization:
						`Bearer ${jwtKey}`,
				},
			})
			.then((response) => {
				console.log(response);
			})
			.catch((error) => {
				console.log(error.toString());
			});
	};
	return (
		<Navbar>
			<VStack alignItems="start" w="full">
				<HStack mb={{ base: 4, lg: 6 }}>
					<VStack alignItems="start">
						<Text textStyle="title">Create campaigns</Text>
						<Text textStyle="subtitle">Create Campaigns</Text>
					</VStack>
				</HStack>
				<HStack>
					<VStack>
						<form onSubmit={handleSubmit}>
							<FormControl as="fieldset">
								<FormLabel mt={4}>Campaign Name</FormLabel>

								<Input
									placeholder="Campaign Name"
									onChange={(event) => {
										setCampaignName(event.currentTarget.value);
									}}
								/>
								<FormLabel mt={4}>Campaign Program</FormLabel>
								<RadioGroup
									defaultValue="Shopping"
									onChange={(event) => {
										setRewardProgram(event);
									}}
								>
									<HStack spacing="24px">
										<Radio value="Shopping">Shopping</Radio>
										<Radio value="Premium Miles">Premium Miles</Radio>
										<Radio value="Platinum Miles">Platinum Miles</Radio>
										<Radio value="Freedom">Freedom</Radio>
									</HStack>
								</RadioGroup>

								<FormLabel mt={4}>Amount</FormLabel>
								<NumberInput
									max={50}
									min={1}
									onChange={(event) => {
										setRewardAmount(event);
									}}
								>
									<NumberInputField />
									<NumberInputStepper>
										<NumberIncrementStepper />
										<NumberDecrementStepper />
									</NumberInputStepper>
								</NumberInput>
								<FormLabel mt={4}>Minimum Spend Amount</FormLabel>
								<NumberInput
									min={1}
									onChange={(event) => {
										setMinSpend(event);
									}}
								>
									<NumberInputField />
									<NumberInputStepper>
										<NumberIncrementStepper />
										<NumberDecrementStepper />
									</NumberInputStepper>
								</NumberInput>
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

									<FormLabel mt={4}>End Date</FormLabel>
									<Input
										placeholder="Select Date and Time"
										size="md"
										type="datetime-local"
										onChange={(event) => {
											setEndDate(event.currentTarget.value);
										}}
									/>
								</FormControl>

								<FormLabel mt={4}>Applicable Merchant</FormLabel>
								<Input
									placeholder=""
									onChange={(event) => {
										setMerchant(event.currentTarget.value);
									}}
								/>

								<FormLabel mt={4}>Applicable MCC</FormLabel>
								<NumberInput
									max={9999}
									min={1}
									onChange={(event) => {
										setMcc(event.currentTarget.value);
									}}
								>
									<NumberInputField />
									<NumberInputStepper>
										<NumberIncrementStepper />
										<NumberDecrementStepper />
									</NumberInputStepper>
								</NumberInput>

								<FormLabel mt={4}>For Foreign Currency</FormLabel>
								<RadioGroup
									// defaultValue={false}
									onChange={(event) => {
										setForeignCurrency(event === "true");
									}}
								>
									<HStack spacing="24px">
										<Radio value={"false"}>Local Transactions Only</Radio>
										<Radio value={"true"}>Applicable for both</Radio>
									</HStack>
								</RadioGroup>
								<Button
									variantColor="teal"
									variant="outline"
									type="submit"
									width="full"
									mt={4}
									onClick = {()=>addCampaign()}
								>
									Add Campaign
								</Button>
							</FormControl>
						</form>
					</VStack>
				</HStack>
			</VStack>
		</Navbar>
	);
}
