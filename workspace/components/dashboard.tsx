import { Box, Container, Grid, GridItem, Heading, Input, Stack, Switch, Table, TableCaption, TableContainer, Tbody, Td, Text, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Filler,
  Legend,
} from 'chart.js'

import { Line } from 'react-chartjs-2'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Filler,
  Legend
);

export const options = {
  responsive: true,
  plugins: {
    legend: {
      display: false,
    },
    title: {
      display: false,
    },
  },
};

const labels = ['1h', '2h', '3h', '4h', '5h', '6h', '7h', '8h', '9h', '10h'];

export const data = {
  labels,
  datasets: [
    {
      fill: true,
      label: 'Hoje',
      data: [68, 56, 24, 40, 34, 35, 46, 23, 56, 23],
      borderColor: 'rgb(255, 255, 255)',
      backgroundColor: 'rgba(255, 255, 255, 0.1)',
    },
    {
      fill: true,
      label: 'Ontem',
      data: [12, 32, 65, 21, 23, 44, 12, 32, 12, 54],
      borderColor: 'rgb(82, 82, 82)',
      backgroundColor: 'rgba(82, 82, 82, 0.1)',
    },
  ],
};

export default function Dashboard({ watcher }: any) {
  return (
    <Box bg='#121212' minH={'100vh'}>
      <Box bg={'crimson'} h={3} />

      <Container maxW='1280px' textColor={'white'} h={'100%'} pb={7}>
        <Grid
          templateColumns='repeat(3, 1fr)'
          gap={4}
        >
          <GridItem colSpan={4} bg='#1f1e1e' mt={6} p={7}>
            <Text>Filtros aqui</Text>
          </GridItem>

          <GridItem rowSpan={4} colSpan={1} bg={'#1f1e1e'} p={7}>
            <Heading fontSize={'24px'} mb={3}>{watcher.name}</Heading>
            <Text textColor={'gray'}>{watcher.desc}</Text>

            <Input placeholder='Search routine' mt={3} variant='unstyled' size='lg' />
            <Box borderBottom={'1px solid #262525'} mt={4} />

            <Box overflow={'auto'} maxH='550px'>
              {[1, 2, 3, 4, 5].map((variant) => (
                <Box
                  key={variant}
                  borderBottom={'1px solid #262525'}
                  p={4}
                  _hover={{
                    bg: '#0f0f0f',
                    cursor: 'pointer'
                  }}
                >
                  <Heading fontSize={'17px'} mb={3}>📟 Rotina {variant}</Heading>
                  <Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</Text>
                </Box>
              ))}
            </Box>

          </GridItem>

          <GridItem colSpan={3} bg='#1f1e1e' p={7}>
            <Stack align='center' justifyContent={'space-between'} direction='row'>
              <Heading fontSize={'24px'} mb={3}>🌵 Rotina 01</Heading>
              <Switch size='md' colorScheme={'orange'} />
            </Stack>
            <Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</Text>
          </GridItem>

          <GridItem colSpan={3}>
            <Grid templateColumns='repeat(3, 1fr)' gap={4}>
              <GridItem colSpan={1} bg='#1f1e1e' p={7}>
                <Heading fontSize={'54px'} mb={3}>1.56ms</Heading>
                <Text>Indicator 01</Text>
                <Text textColor={'gray'}>Lorem Ipsum is simply</Text>
              </GridItem>
              <GridItem colSpan={1} bg='#1f1e1e' p={7}>
                <Heading fontSize={'54px'} mb={3}>23%</Heading>
                <Text>Indicator 01</Text>
                <Text textColor={'gray'}>Lorem Ipsum is simply</Text>
              </GridItem>
              <GridItem colSpan={1} bg='#1f1e1e' p={7}>
                <Heading fontSize={'54px'} mb={3}>0.21</Heading>
                <Text>Indicator 01</Text>
                <Text textColor={'gray'}>Lorem Ipsum is simply</Text>
              </GridItem>
            </Grid>
          </GridItem>

          <GridItem colSpan={3} bg='#1f1e1e' p={7}>
            <Text>Grafico geral com todas as rotinas aqui</Text>
            <Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</Text>
            <Box mt={5}>
              <Line options={options} data={data} height={100} />
            </Box>
          </GridItem>

          <GridItem colSpan={2} bg='#1f1e1e' p={7}>
            <TableContainer>
              <Table variant='simple'>
                <TableCaption>Imperial to metric conversion factors</TableCaption>
                <Thead>
                  <Tr>
                    <Th>To convert</Th>
                    <Th>into</Th>
                    <Th isNumeric>multiply by</Th>
                  </Tr>
                </Thead>
                <Tbody>
                  <Tr>
                    <Td>inches</Td>
                    <Td>millimetres (mm)</Td>
                    <Td isNumeric>25.4</Td>
                  </Tr>
                  <Tr>
                    <Td>feet</Td>
                    <Td>centimetres (cm)</Td>
                    <Td isNumeric>30.48</Td>
                  </Tr>
                  <Tr>
                    <Td>yards</Td>
                    <Td>metres (m)</Td>
                    <Td isNumeric>0.91444</Td>
                  </Tr>
                </Tbody>
                <Tfoot>
                  <Tr>
                    <Th>To convert</Th>
                    <Th>into</Th>
                    <Th isNumeric>multiply by</Th>
                  </Tr>
                </Tfoot>
              </Table>
            </TableContainer>
          </GridItem>
        </Grid>
        <GridItem colSpan={2} py={7}>
          <Text>Outis Dashboard - v0.0.1</Text>
        </GridItem>
      </Container>
    </Box>
  )
}
