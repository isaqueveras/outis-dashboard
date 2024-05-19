import { Badge, Box, Container, Heading, Link, Table, Tbody, Td, Text, Th, Thead, Tr } from "@chakra-ui/react";

export default function LogPage() {
  return (
    <Box bg='#121212' minH={'100vh'}>
      <Box bg={'crimson'} h={2} />
      <Container maxW='8xl' textColor={'white'} my={7}>
        <Box
          display={'flex'}
          justifyContent={'space-between'}
          alignItems={'end'}
          borderBottom={'1px solid #262525'}
          pb={4}
        >
          <Box>
            <Heading mt={2}>Technology image collector</Heading>
            <Text textColor={'gray'}>Lorem Ipsum is simply dummy...</Text>
          </Box>
          <Box textAlign={'right'}>
            <Text>19/05/2024 02:06</Text>
            <Text>
              <Badge variant='subtle' px={2} py={1} colorScheme='gray'>
                ID: {' '}
                <Link href="/w/3ed4bbc9-ab67-402f-82c1-0fdbd004afd6/r/5b15145b-df20-478d-bcc1-9dd30a451ad0/d/f6881fdd-3b4a-4b3a-a43d-22eabbb3b146">
                  305b4170-8421-4fd2-8066-8be1dcbc0159
                </Link>
              </Badge>
            </Text>
          </Box>
        </Box>

        <Table variant={'striped'} colorScheme={'blackAlpha'} mt={4}>
          <Thead>
            <Tr>
              <Th textColor={'white'}>Type</Th>
              <Th textColor={'white'}>Message</Th>
              <Th textColor={'white'}>Timestamp</Th>
            </Tr>
          </Thead>
          <Tbody>
            <Tr>
              <Td>INFO</Td>
              <Td>this is an information message</Td>
              <Td>2024-05-05T14:47:16.084211061-03:00</Td>
            </Tr>
            <Tr>
              <Td>ERROR</Td>
              <Td>não consegui desbloquear a rotina :/</Td>
              <Td>2024-05-05T14:47:16.084242814-03:00</Td>
            </Tr>
            <Tr>
              <Td>DEBUG</Td>
              <Td>Está funcionando aqui</Td>
              <Td>2024-05-05T14:47:16.084239624-03:00</Td>
            </Tr>
            <Tr>
              <Td>ERROR</Td>
              <Td>não consegui desbloquear a rotina :/</Td>
              <Td>2024-05-05T14:47:16.084242814-03:00</Td>
            </Tr>
            <Tr>
              <Td>ERROR</Td>
              <Td>There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour.</Td>
              <Td>2024-05-05T14:47:16.084242814-03:00</Td>
            </Tr>
            <Tr>
              <Td>ERROR</Td>
              <Td>não consegui desbloquear a rotina :/</Td>
              <Td>2024-05-05T14:47:16.084242814-03:00</Td>
            </Tr>
          </Tbody>
        </Table>
      </Container>
    </Box>
  )
}
