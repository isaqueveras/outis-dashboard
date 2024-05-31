import { useFetch } from "@/app/hooks/useFetch"
import { Badge, Box, Heading, Skeleton, Stack, Text } from "@chakra-ui/react"

const RoutineDetails = ({ isLoading }: any) => {
  const { data, error, isValidating } = useFetch('/api/watcher/84cbfa32/routine/84cbfa31')
  if (error) return <div>falhou ao carregar</div>

  let state: string = data?.state == "stoped" ? 'gray' : 'blue'

  return (
    <Stack align='center' justifyContent={'space-between'} alignItems={'start'} direction='row' p={7}>
      <Box>
        <Skeleton mb={1} isLoaded={!isValidating || !isLoading}>
          <Heading fontSize={'20px'} mb={1}>{data?.name}</Heading>
        </Skeleton>
        <Text textColor={'gray'}>{data?.desc}</Text>
      </Box>
      <Box>
        {/* <Switch size='md' colorScheme={'green'} isChecked={true} /> */}
        <Badge variant='subtle' px={2} py={1} colorScheme='gray'>ID: 3ed4bbc9-ab67-402f-82c1-0fdbd004afd6</Badge>
        <Badge variant='solid' ml={1} px={3} py={1} colorScheme={state}>{data?.state}</Badge>
      </Box>
    </Stack>
  )
}

export default RoutineDetails
