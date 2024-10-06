-- DEFINE TABLE tenant SCHEMALESS;
-- DEFINE TABLE permission SCHEMALESS;
-- DEFINE TABLE role SCHEMALESS;

-- CREATE permission:create SET name = 'create content';
-- CREATE permission:read SET name = 'read content';
-- CREATE permission:update SET name = 'update content';
-- CREATE permission:delete SET name = 'delete content';
-- CREATE role:reader SET name = 'Content Reader';

-- RELATE role:reader->role_permission->permission:read;

-- CREATE role:author SET name = 'Content Author';

-- RELATE role:author->role_permission->permission:read;
-- RELATE role:author->role_permission->permission:create;
-- RELATE role:author->role_permission->permission:update;

-- CREATE role:admin SET name = 'Administrator';

-- RELATE role:admin->role_permission->permission:read;
-- RELATE role:admin->role_permission->permission:create;
-- RELATE role:admin->role_permission->permission:update;
-- RELATE role:admin->role_permission->permission:delete;

-- -- create the tenant for our cat owners
-- CREATE tenant:cat SET name = 'Cat Owners';

-- RELATE role:admin->tenant_role->tenant:cat;
-- RELATE role:author->tenant_role->tenant:cat;
-- RELATE role:reader->tenant_role->tenant:cat;

-- -- create the tenant for our musicians
-- CREATE tenant:musician SET name = 'Musicians';

-- RELATE role:admin->tenant_role->tenant:musician;
-- RELATE role:author->tenant_role->tenant:musician;
-- RELATE role:reader->tenant_role->tenant:musician;

-- -- create the user
-- LET $createdUser = (CREATE user:1 SET
--   name='John Doe',
--   firstName = 'John',
--   lastName = 'Doe',
--   email='john.doe@example.com');

-- -- link the user John Doe to the tenant "car enthusiasts"
-- LET $memberRelation = (RELATE $createdUser->tenant_member->tenant:car);

-- -- link the member relationship to a role
-- RELATE $memberRelation->member_role->role:author;

-- -- finally, return the newly created user
-- RETURN $createdUser;