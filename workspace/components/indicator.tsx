import { Box, Heading, Text } from "@chakra-ui/react"

type IndicatorProps = {
  name: string
  desc: string
  value: any
}

const Indicator = (indicator: IndicatorProps) => {
  return (
    <Box>
      <Heading fontSize={'54px'} mb={3}>{indicator.value}</Heading>
      <Text>{indicator.name}</Text>
      <Text textColor={'gray'}>{indicator.desc}</Text>
    </Box>
  )
}

export default Indicator
