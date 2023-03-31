import { Box, Spinner } from '@chakra-ui/react';
import Image from 'next/image';
import React, { useState } from 'react';
import Marquee from 'react-fast-marquee';

function CustomMarquee() {
    var full = [...Array(33).keys()]
    full = [full.splice(0, 8), full.splice(0, 8), full.splice(0, 8), full.splice(0, 10)]

    return (
        <Box maxW="100%" w="100%">
            {full.map((array, index) => {
                return (
                    <Marquee
                        key={index}
                        gradientColor={false}
                        direction={index % 2 ? 'right' : 'left'}
                    >
                        {array.map((num) => {
                            return (
                                <Box mx={4} my={3} key={num}>
                                    <img id={`/merchant${num}.webp`} src={`https://ik.imagekit.io/alvinowyong/g1t2/merchant${num}.webp?tr=w-60`} width="0" height="0" sizes="100vw" style={{ width: '35px', height: 'auto' }} />
                                </Box>
                            )
                        })}
                    </Marquee>
                )

            })}
        </Box>

    );
}

export default CustomMarquee;