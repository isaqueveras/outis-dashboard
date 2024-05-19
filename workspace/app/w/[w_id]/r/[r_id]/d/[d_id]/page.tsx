import { Box, Code, Container, Heading, Text } from "@chakra-ui/react";

export default function DetailsPage() {
  return (
    <Box bg='#121212' minH={'100vh'}>
      <Box bg={'crimson'} h={2} />
      <Container maxW='8xl' textColor={'white'} my={7}>
        <Heading my={2}>Details Page</Heading>
        <Code>
          {`"client": {
          "address": {
          "id": "fe4876a3-685a-41e9-8d8a-9ae1bd4a288d",
        "number": "S/N",
        "street": "Av. 01"
      },
        "email": "antonio.francisco.silva@email.com",
        "id": "a46de7c6-2264-4a0f-98e3-3fe579b14d26",
        "name": "Antonio Francisco da Silva"
    },
        "uuid": [
        "2ae3419e-cf31-4dbe-af06-29c80783ee52",
        "03eb9570-83e1-467e-bcc0-6608f887ad22",
        "42c2edf7-3b29-4259-899a-f0139fc6dea4"
        ]`}
        </Code>
      </Container>
    </Box>
  )
}
