import {
    Button,
    Checkbox,
    FormControl,
    FormLabel,
    Input,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    NumberInput,
    NumberInputField,
    Radio,
    RadioGroup,
    Stack,
} from '@chakra-ui/react';
import axios from 'axios';
import { useRef, useState } from 'react';

export default function AddCampaigns(props) {
    const [campaignName, setCampaignName] = useState("");
    const [minSpend, setMinSpend] = useState(0);
    const [startDate, setStartDate] = useState(new Date());
    const [endDate, setEndDate] = useState(new Date());
    const [rewardProgram, setRewardProgram] = useState("Shopping");
    const [rewardAmount, setRewardAmount] = useState(0);
    const [mcc, setMcc] = useState("0");
    const [foreignCurrency, setForeignCurrency] = useState(false);
    const [merchant, setMerchant] = useState("");
    const base_reward = useRef();


    const addCampaign = () => {
        props.toast({
            title: "In progress",
            description: "Please hold on while we create a campaign",
            status: "info",
            duration: 9000,
            isClosable: true,
        });
        const body = {
            name: campaignName,
            min_spend: parseFloat(minSpend),
            base_reward: base_reward.current.checked,
            start_date: startDate + ":00Z",
            end_date: endDate + ":00Z",
            reward_program: rewardProgram,
            reward_amount: parseInt(rewardAmount),
            mcc: mcc.toString(),
            merchant: merchant,
            foreign_currency: foreignCurrency,
        };

        if (!Object.values(body).every(value => value)) {
            props.toast.closeAll();
            props.toast({
                title: "Empty field(s)",
                description: "Please fill in all the fields",
                status: "warning",
                duration: 9000,
                isClosable: true,
            });
            return
        }
        axios
            .post(`https://itsag1t2.com/campaign`, body, {
                headers: {
                    Authorization: props.session.id,
                },
            })
            .then((response) => {
                props.toast.closeAll();
                props.toast({
                    title: "Success",
                    description: `Campaign created successfully`,
                    status: "success",
                    duration: 9000,
                    isClosable: true,
                });
                props.refresh()
                props.onClose()
            })
            .catch((error) => {
                console.log(error);
                props.toast.closeAll();
                props.toast({
                    title: "Error",
                    description: "An error occurred while creating a campaign",
                    status: "error",
                    duration: 9000,
                    isClosable: true,
                });
            });
    };
    return (
        <>
            <Modal isOpen={props.isOpen} onClose={props.onClose} size={{ base: "md", md: "xl" }}>
                <ModalOverlay />
                <ModalContent >
                    <ModalHeader fontSize="md">Add campaign</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <FormControl as="fieldset" >
                            <FormLabel mt={4} fontSize={{ base: "small", md: "sm" }} fontWeight={600}>Campaign name</FormLabel>
                            <Input
                                fontSize="small"
                                placeholder="e.g. 10% off on shopping on Ascenda"
                                onChange={(event) => {
                                    setCampaignName(event.currentTarget.value);
                                }}
                            />
                            <FormLabel mt={4} fontSize={{ base: "small", md: "sm" }} fontWeight={600}>Campaign program</FormLabel>
                            <RadioGroup
                                defaultValue="Shopping"
                                onChange={(event) => {
                                    setRewardProgram(event);
                                }}
                            >
                                <Stack spacing={{ base: 0, md: "24px" }} fontSize={{ base: "small", md: "md" }} direction={{ base: "column", md: "row" }}>
                                    <Radio size="sm" value="scis_shopping">Shopping</Radio>
                                    <Radio size="sm" value="scis_premiummiles">PremiumMiles</Radio>
                                    <Radio size="sm" value="scis_platinummiles">PlatinumMiles</Radio>
                                    <Radio size="sm" value="scis_freedom">Freedom</Radio>
                                </Stack>
                            </RadioGroup>

                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>Reward amount (% OR reward/$)</FormLabel>
                            <NumberInput
                                max={50}
                                min={1}
                                onChange={(event) => {
                                    setRewardAmount(event);
                                }}
                            >
                                <NumberInputField fontSize="small" />
                            </NumberInput>
                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>Minimum Spend Amount</FormLabel>
                            <NumberInput
                                min={1}
                                onChange={(event) => {
                                    setMinSpend(event);
                                }}
                            >
                                <NumberInputField fontSize="small" />
                            </NumberInput>
                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>Base reward</FormLabel>
                            <Checkbox size="sm" ref={base_reward}>Yes, applicable</Checkbox>
                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>For Foreign Currency</FormLabel>
                            <RadioGroup
                                onChange={(event) => {
                                    setForeignCurrency(event === "true");
                                }}
                            >
                                <Stack spacing={{ base: 0, md: "24px" }} fontSize={{ base: "small", md: "md" }} direction={{ base: "column", md: "row" }}>
                                    <Radio size="sm" value={"false"}>Local Transactions Only</Radio>
                                    <Radio size="sm" value={"true"}>Applicable for both</Radio>
                                </Stack>
                            </RadioGroup>
                            <FormControl isRequired>
                                <FormLabel mt={4} fontSize="sm" fontWeight={600}>Start Date</FormLabel>
                                <Input
                                    fontSize="small"
                                    placeholder="Select Date and Time"
                                    size="md"
                                    type="datetime-local"
                                    onChange={(event) => {
                                        setStartDate(event.currentTarget.value);
                                    }}
                                />

                                <FormLabel mt={4} fontSize="sm" fontWeight={600}>End Date</FormLabel>
                                <Input
                                    fontSize="small"
                                    placeholder="Select Date and Time"
                                    size="md"
                                    type="datetime-local"
                                    onChange={(event) => {
                                        setEndDate(event.currentTarget.value);
                                    }}
                                />
                            </FormControl>

                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>Applicable Merchant</FormLabel>
                            <Input
                                fontSize="small"
                                placeholder=""
                                onChange={(event) => {
                                    setMerchant(event.currentTarget.value);
                                }}
                            />

                            <FormLabel mt={4} fontSize="sm" fontWeight={600}>Applicable MCC</FormLabel>
                            <NumberInput
                                max={9999}
                                min={1}
                                defaultValue={0}
                                onChange={(event) => {
                                    setMcc(event);
                                }}
                            >
                                <NumberInputField fontSize="small" />
                            </NumberInput>


                        </FormControl>
                    </ModalBody>

                    <ModalFooter>
                        <Button size="sm" variant='ghost' colorScheme='purple' mr={3} onClick={props.onClose}>
                            Close
                        </Button>
                        <Button size="sm" colorScheme='purple' onClick={() => addCampaign()}>Submit</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
}
