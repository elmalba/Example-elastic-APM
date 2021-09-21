import React,{useState,useEffect} from 'react';
import { Table, Input } from 'antd';

import apm from './apm'


const { Search } = Input;

  const columns = [
    {
      title: 'Nombre Jugador',
      dataIndex: 'FullName',
      key: 'FullName',
    },
    {
      title: 'Puntaje',
      dataIndex: 'Point',
      key: 'Point',
    }
  ];
  
export default function Ranking (){
    const [players,setPlayers] = useState([])
    const [query, setQuery] = useState("")
    useEffect(() => {
        const transaction = apm.startTransaction('Click get Data', 'custom')
        const url = '/api/ranking?q=' + query
        const httpSpan = transaction.startSpan('GET ' + url, 'external.http')
        fetch(url)
        .then(response =>{
         if (!response.ok) {
              apm.captureError(new Error(`fetch failed with status ${response.status} ${response.statusText}`))
         }
         httpSpan.end()
         transaction.end()


        return  response.json()
        })
        .then(data => setPlayers(data));
    }, [query])
    let searchQuery=(value) => {
        setQuery(value)
    }

    return  <div>
         <div style={{ marginBottom: 16 }}>
         <Search addonBefore="Nombre del jugador"  placeholder="Ingrese el nombre del jugardor" onSearch={searchQuery} enterButton="Buscar"  />

        </div>
        <Table dataSource={players} columns={columns} />

        </div>
  }
