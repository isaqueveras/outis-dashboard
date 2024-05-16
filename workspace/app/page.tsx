'use client'

import {
	Box,
	Container,
	Grid,
	GridItem,
	Heading,
	Input,
	Stack,
	Switch,
	Table,
	TableCaption,
	TableContainer,
	Tbody,
	Td,
	Text,
	Tfoot,
	Th,
	Thead,
	Tr
} from '@chakra-ui/react'

export default function Home() {
	return (
		<Box bg='#121212' minH={'100vh'}>
			<Box bg={'crimson'} h={3} />

			<Container maxW='1300px' textColor={'white'} h={'100%'} pb={7}>
				<Grid
					// templateRows='repeat(4, 1fr)'
					templateColumns='repeat(3, 1fr)'
					gap={4}
				// h={'100%'}
				>
					<GridItem colSpan={4} bg='#1f1e1e' mt={6} p={7}>
						<Text>Filtros aqui</Text>
					</GridItem>

					<GridItem rowSpan={3} colSpan={1} bg={'#1f1e1e'} p={7}>
						<Heading fontSize={'24px'} mb={3}>🎫 Minhas Rotinas</Heading>
						<Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry.</Text>

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

					<GridItem colSpan={1} bg='#1f1e1e' p={7}>
						<Text>Grafico geral com todas as rotinas aqui</Text>
						<Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s</Text>
					</GridItem>

					<GridItem colSpan={1} bg='#1f1e1e' p={7}>
						<Text>Indicador geral das rotinas</Text>
						<Text textColor={'gray'}>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s</Text>
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
	);
}
