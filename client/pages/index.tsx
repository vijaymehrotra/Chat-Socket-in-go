import { useState, useEffect, useContext } from 'react'
import { API_URL } from '../constants'
import { v4 as uuidv4 } from 'uuid'
import { WEBSOCKET_URL } from '../constants'
import { AuthContext } from '../modules/auth_provider'
import { WebsocketContext } from '../modules/websocket_provider'
import { useRouter } from 'next/router'

const index = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([])
  const [roomName, setRoomName] = useState('')
  const { user } = useContext(AuthContext)
  const { setConn } = useContext(WebsocketContext)

  const router = useRouter()

  const getRooms = async () => {
    try {
      const res = await fetch(`${API_URL}/ws/get-rooms`, {
        method: 'GET',
      })

      const data = await res.json()
      console.log('Rooms data:', data) // Debug: Check the API response

      if (res.ok) {
        setRooms(data.rooms)
      } else {
        console.error('Unexpected response format:', data)
        setRooms([]) // Fallback in case data is not an array
      }
    } catch (err) {
      console.log('Error fetching rooms:', err)
      setRooms([]) // Fallback in case of error
    }
  }

  useEffect(() => {
    getRooms()
  }, [])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    try {
      setRoomName('')
      const res = await fetch(`${API_URL}/ws/create-room`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({
          id: uuidv4(),
          name: roomName,
        }),
      })

      if (res.ok) {
        getRooms()
      }
    } catch (err) {
      console.log('Error creating room:', err)
    }
  }

  const joinRoom = (roomId: string) => {
    if (!user?.id || !user?.username) {
      console.log(user)
      console.error('User ID or username is undefined');
      return;
    }
  
    const ws = new WebSocket(
      `${WEBSOCKET_URL}/ws/join-room/${roomId}?user_id=${user.id}&username=${user.username}`
    );
  
    ws.onopen = () => {
      setConn(ws);
      router.push('/app');
    };
  
    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };
  };

  return (
    <>
      <div className='my-8 px-4 md:mx-32 w-full h-full'>
        <div className='flex justify-center mt-3 p-5'>
          <input
            type='text'
            className='border border-grey p-2 rounded-md focus:outline-none focus:border-blue'
            placeholder='room name'
            value={roomName}
            onChange={(e) => setRoomName(e.target.value)}
          />
          <button
            className='bg-blue border text-white rounded-md p-2 md:ml-4'
            onClick={submitHandler}
          >
            create room
          </button>
        </div>
        <div className='mt-6'>
          <div className='font-bold'>Available Rooms</div>
          <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
            {rooms?.length > 0 ? (
              rooms.map((room, index) => (
                <div
                  key={index}
                  className='border border-blue p-4 flex items-center rounded-md w-full'
                >
                  <div className='w-full'>
                    <div className='text-sm'>room</div>
                    <div className='text-blue font-bold text-lg'>{room.name}</div>
                  </div>
                  <div className=''>
                    <button
                      className='px-4 text-white bg-blue rounded-md'
                      onClick={() => joinRoom(room.id)}
                    >
                      join
                    </button>
                  </div>
                </div>
              ))
            ) : (
              <div>No rooms available</div> // Display this if no rooms are available
            )}
          </div>
        </div>
      </div>
    </>
  )
}

export default index
