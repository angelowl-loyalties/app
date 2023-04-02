import { HStack, Stat, StatGroup, StatHelpText, StatLabel, StatNumber, Text, VStack } from '@chakra-ui/react';
import axios from 'axios';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

import Navbar from '../components/Navbar';
import Loading from '../components/Loading';

export default function Home() {
    const [loading, setLoading] = useState(true)
    const [cards, setCards] = useState({})
    const router = useRouter()
    const types = ['scis_shopping', 'scis_freedom', 'scis_platinummiles', 'scis_premiummiles']

    const { data: session, status } = useSession({
        required: true,
        onUnauthenticated() {
            router.push("/login")
        },
    });

    useEffect(() => {
        if (!session) {
            return
        }
        axios.get(`https://itsag1t2.com/user/${session.userId}`, { headers: { Authorization: session.id } })
            .then((response) => {
                setCards(response.data.data.CreditCards)
                response.data.data.CreditCards.map((card, index) => {
                    console.log(index)
                    axios.get(`https://itsag1t2.com/reward/total/${card.id}`, { headers: { Authorization: session.id } })
                        .then((response) => {
                            console.log(cards[index])
                            cards[index].total = response.data.data

                            axios.get(`https://itsag1t2.com/card/type?${card.card_type}`, { headers: { Authorization: session.id } })
                            .then((response) => {
                                cards[index].card_type = response.data.data[0]
                                setLoading(false)
                            }).catch((error) => {
                                setLoading(false)
                                console.log(error)
                            })
                            
                        }).catch((error) => {
                            setLoading(false)
                            console.log(error)
                        })
                    console.log(cards)
                })
            })
    }, [session])


    return (
        <>
            {loading ? <Loading /> :
                <Navbar user>
                    <HStack mb={{ base: 4, lg: 6 }}>
                        <VStack alignItems='start'>
                            <Text textStyle="title">Dashboard</Text>
                            <Text textStyle="subtitle">
                                Supercharge your everyday credit cards and get rewarded when you spend
                            </Text>
                            <HStack spacing={28}>
                                {types.map((type) => {
                                    const card = cards.filter((card) => {
                                        // console.log(card)
                                    })
                                    if (card.length === 0) {
                                        return
                                    }
                                    const total = cards.reduce((total, card) => total + card.total, 0)
                                    return (
                                        <StatGroup key={type}>
                                            <Stat>
                                                <StatLabel>{card[0].card_type.reward_unit}</StatLabel>
                                                <StatNumber>{total}</StatNumber>
                                                <StatHelpText>
                                                    23.36%
                                                </StatHelpText>
                                            </Stat>
                                        </StatGroup>
                                    )
                                })}
                            </HStack>
                        </VStack>
                    </HStack>
                </Navbar>
            }
        </>
    );
}

