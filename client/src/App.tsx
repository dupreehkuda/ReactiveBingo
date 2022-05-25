import { useState } from 'react'
import useSWR from 'swr'
import './App.css'
import { Center, Space, Group, TextInput, Button, SimpleGrid, Stack } from '@mantine/core';

export const ENDPOINT = 'http://localhost:4000'

const fetcher = (url: string) => 
  fetch(`${ENDPOINT}/${url}`).then(r => r.json());

let userNumbers = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]

function App() {
  const {data, mutate} = useSWR<[]>('api/numbers', fetcher);
  var [value, setValue] = useState('');
  var [userBingoPoints, setPoints] = useState(0);

  async function bingoCheck(array: number[]) {
    console.log(JSON.stringify(array))
    const result = await fetch(`${ENDPOINT}/api/bingocheck`, {
      method: 'POST',
      headers: {
          "Content-Type": "application/json"
      },
      body: JSON.stringify(array),
    }).then((r) => r.json());
    setPoints(result)
  }

  function submited(num: number) {
    if (num <= 16 && data != null) {
      var arr:number[] = data
      userNumbers[arr.indexOf(num)] = num
    }
    bingoCheck(userNumbers)
    setValue(value = '')
  }

  function correct(num: number) {
    if (num == 0) {
      return ""
    } else {
      return num
    }
  }

  return (
    <div className="App">
      <Group position='center'>
      <Stack>
      <p className='Title_thing'>{`Bingo counter: ${userBingoPoints}`}</p>
      <Center>
      <SimpleGrid cols={4} spacing='xl'>
        {userNumbers.map(num => ( 
          <div className={`${num != 0 ? "rcorners2" : "rcorners1"}`}>
            <p className='Box_content'>{correct(num)}</p>
          </div>
        ))}
      </SimpleGrid>
      </Center>

      <Space h="xl" />

      <Group position='center'> 
        <TextInput
          placeholder="Done task index"
          variant="filled"
          size="md"
          value={value} 
          onChange={(event) => setValue(event.currentTarget.value)}/>
        <Button onClick={() => submited(Number(value))}
        color="red" 
        size="md">
          Submit
        </Button>
      </Group>
      </Stack>
      <Space h="xl"/>
      </Group>
    </div>
  )
}

export default App