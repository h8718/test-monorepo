import request from 'supertest'

import common from "../common"

const namespaceName = "callpathtest"


describe('Test subflow behaviour', () => {
    beforeAll(common.helpers.deleteAllNamespaces)
    //afterAll(common.helpers.deleteAllNamespaces)

    it(`should create a namespace`, async () => {
        var req = await request(common.config.getDirektivHost()).put(`/api/namespaces/${namespaceName}`)

        expect(req.statusCode).toEqual(200)
        expect(req.body).toMatchObject({
            namespace: {
                createdAt: expect.stringMatching(common.regex.timestampRegex),
                updatedAt: expect.stringMatching(common.regex.timestampRegex),
                name: namespaceName,
                oid: expect.stringMatching(common.regex.uuidRegex),
            },
        })
    })

    it(`should create a directory called /a`, async () => {
        var createDirectoryResponse = await request(common.config.getDirektivHost()).put(`/api/namespaces/${namespaceName}/tree/a?op=create-directory`)
        expect(createDirectoryResponse.statusCode).toEqual(200)
    })

    it(`should create a workflow called /a/child.yaml`, async () => {

        const res = await request(common.config.getDirektivHost())
            .put(`/api/namespaces/${namespaceName}/tree/a/child.yaml?op=create-workflow`)
            .set({
                'Content-Type': 'text/plain',
            })
            .send(`
states:
- id: a
  type: noop
  transform:
    result: 'jq(.input + 1)'`)

        expect(res.statusCode).toEqual(200)
        expect(res.body).toMatchObject({
            namespace: namespaceName,
        })
    })

    it(`should create a workflow called /a/parent1.yaml`, async () => {

        const res = await request(common.config.getDirektivHost())
            .put(`/api/namespaces/${namespaceName}/tree/a/parent1.yaml?op=create-workflow`)
            .set({
                'Content-Type': 'text/plain',
            })
            .send(`
functions:
- id: child
  type: subflow
  workflow: '/a/child.yaml'
states:
- id: a
  type: action
  action:
    function: child
    input: 
      input: 1
  transform:
    result: 'jq(.return.result)'
`)

        expect(res.statusCode).toEqual(200)
        expect(res.body).toMatchObject({
            namespace: namespaceName,
        })
    })

    it(`should invoke the '/a/parent1.yaml' workflow`, async () => {
        const req = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/tree/a/parent1.yaml?op=wait`)
        expect(req.statusCode).toEqual(200)
        expect(req.body).toMatchObject({
            result: 2,
        })
    })

    it(`check if child logs are present in parent's log view`, async () => {
        var instances = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/instances`)
        expect(instances.statusCode).toEqual(200)
        expect(instances.body.instances.results.length).not.toBeLessThan(1)
    })


    it(`check if parentslogs contain child logs`, async () => {
        var instancesSource = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/instances?filter.field=AS&filter.type=WORKFLOW&filter.val=a/parent1.yaml`)
        var instancesChild = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/instances?filter.field=AS&filter.type=WORKFLOW&filter.val=/a/child.yaml`)
        // create a new array containing the ids
        const idsS = instancesSource.body.instances.results.map((result) => result.id)
        const idsC = instancesChild.body.instances.results.map((result) => result.id)
        const logsSourceResponse = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/instances/${idsS}/logs`)
        expect(logsSourceResponse.statusCode).toEqual(200)
        const logsChildResponse = await request(common.config.getDirektivHost()).get(`/api/namespaces/${namespaceName}/instances/${idsC}/logs`)
        expect(logsChildResponse.statusCode).toEqual(200)
        expect(logsChildResponse.body.results.length).toBeLessThan(logsSourceResponse.body.results.length)
        expect(logsChildResponse.body.results.length).not.toBeLessThan(1)

    })

})
