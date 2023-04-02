import {
    Button,
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
} from '@chakra-ui/react';
import axios from 'axios';
import { useState } from 'react';

export default function AddExclusions(props) {
    const [startDate, setStartDate] = useState(new Date());
    const [mcc, setMcc] = useState(0);

    const addExclusion = () => {
        props.toast({
            title: "In progress",
            description: "Please hold on while we create an exclusion",
            status: "info",
            duration: 9000,
            isClosable: true,
        });
        const body = {
            valid_from: startDate + ":00Z",
            mcc: parseInt(mcc),
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
            .post(`https://itsag1t2.com/exclusion`, body, {
                headers: {
                    Authorization: props.session.id,
                },
            })
            .then((response) => {
                console.log(response);
                props.toast.closeAll();
                props.toast({
                    title: "Success",
                    description: `Exclusion created successfully`,
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
                    description: "An error occurred while creating a exclusion",
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
                    <ModalHeader fontSize="md">Add exclusion</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody>
                        <FormControl as="fieldset">
                            <FormControl isRequired>
                                <FormLabel mt={4} fontSize={{ base: "small", md: "sm" }} fontWeight={600}>Start date</FormLabel>
                                <Input
                                    fontSize="small"
                                    type="datetime-local"
                                    onChange={(event) => {
                                        setStartDate(event.currentTarget.value);
                                    }}
                                />
                            </FormControl>

                            <FormLabel mt={4} fontSize={{ base: "small", md: "sm" }} fontWeight={600}>Applicable MCC</FormLabel>
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
                        <Button size="sm" colorScheme='purple' onClick={() => addExclusion()}>Submit</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
}
