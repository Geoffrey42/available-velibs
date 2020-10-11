import React, { useEffect, useState } from 'react';
import { Header, Table } from 'semantic-ui-react';
import axios from 'axios';

const url = "http://localhost:4242/api/fetch"

function ResultsTable() {
    const [records, setRecords] = useState([]);
    const [hits, setHits] = useState(0);
    const [total, setTotal] = useState(0);
    const [distance, setDistance] = useState("0");

    useEffect(() => {
        const interval = setInterval(async () => {
            try {
                const response = await axios(url);

                console.log("response.data: ", response.data);

                setRecords(response.data.records);
                setHits(response.data.nhits);
                setTotal(response.data.total);
                setDistance(response.data.Distance);

            } catch(err) {
                console.error("Could not fetch backend server: ", err);
            }
        }, 1000);
        return () => clearInterval(interval);
    }, [records, hits, total, distance]);

    return (
        <Table>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell>
                        There are {total} available velibs {distance} meters around Splio's HQ at the moment
                    </Table.HeaderCell>
                </Table.Row>
            </Table.Header>
            <Table.Body>
                <Table.Row>
                    <Table.Cell>
                        <Header as='h4'>
                            <Header.Content>
                                [STATION NAME] station
                                <Header.Subheader>
                                    Available velibs: [NUMBER]
                                </Header.Subheader>
                            </Header.Content>
                        </Header>
                    </Table.Cell>
                </Table.Row>
            </Table.Body>
        </Table>
    );
}

export default ResultsTable;