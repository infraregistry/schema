-- -- Remove existing tables
-- REMOVE TABLE component;
-- REMOVE TABLE property;

-- -- Define component table
-- DEFINE TABLE component SCHEMAFULL CHANGEFEED 3d;
-- DEFINE FIELD name ON TABLE component TYPE string;
-- DEFINE FIELD description ON TABLE component TYPE string;
-- DEFINE FIELD status ON TABLE component TYPE string;

-- -- Define property table
-- DEFINE TABLE property SCHEMAFULL CHANGEFEED 3d;
-- DEFINE FIELD name ON TABLE property TYPE string;
-- DEFINE FIELD description ON TABLE property TYPE string;

-- -- Define the relationship
-- DEFINE FIELD component ON TABLE property TYPE record<component>;
-- DEFINE FIELD properties ON TABLE component TYPE array<record<property>>;


-- -- Create a component
-- LET $component = (CREATE component SET name='Component A', description='A sample component', status='active'>;

-- -- Create a property related to the component
-- LET $property = (CREATE property SET name='Property A', description='A sample property', component=$component>;

-- -- Update the component to include the property in the properties array
-- -- UPDATE $component_id SET properties += [$property_id];

-- RELATE $component->properties->$property;


REMOVE TABLE organization;
REMOVE TABLE component;
REMOVE TABLE property;
REMOVE TABLE organization_component;
REMOVE TABLE organization_property;
REMOVE TABLE component_link;
REMOVE TABLE component_property;

-- Organizations (Top Level)

DEFINE TABLE organization SCHEMAFULL CHANGEFEED 3d;
DEFINE FIELD name ON TABLE organization TYPE string;
DEFINE FIELD description ON TABLE organization TYPE string;
DEFINE FIELD status ON TABLE organization TYPE string;
-- DEFINE FIELD created ON TABLE organization TYPE datetime DEFAULT time::now();
DEFINE INDEX name ON TABLE organization COLUMNS name UNIQUE;


DEFINE TABLE organization_component TYPE RELATION;
DEFINE FIELD in ON TABLE organization_component TYPE record<organization>;
DEFINE FIELD out ON TABLE organization_component TYPE record<component>;

-- Component

DEFINE TABLE component SCHEMAFULL CHANGEFEED 3d;
DEFINE FIELD name ON TABLE component TYPE string;
DEFINE FIELD description ON TABLE component TYPE string;
DEFINE FIELD status ON TABLE component TYPE string;

-- Component Link

DEFINE TABLE component_link TYPE RELATION;
DEFINE FIELD in ON TABLE component_link TYPE record<component>;
DEFINE FIELD out ON TABLE component_link TYPE record<component>;
DEFINE FIELD weight ON TABLE component TYPE int DEFAULT 0;

-- Component -> Property

DEFINE TABLE component_property TYPE RELATION;
DEFINE FIELD in ON TABLE component_property TYPE record<component>;
DEFINE FIELD out ON TABLE component_property TYPE record<property>;

-- Component -> Tag

DEFINE TABLE component_tag TYPE RELATION;
DEFINE FIELD in ON TABLE component_tag TYPE record<component>;
DEFINE FIELD out ON TABLE component_tag TYPE record<tag>;

-- Property

DEFINE TABLE property SCHEMAFULL CHANGEFEED 3d;
DEFINE FIELD name ON TABLE property TYPE string;
DEFINE FIELD description ON TABLE property TYPE string;

-- Tag

DEFINE TABLE tag SCHEMAFULL CHANGEFEED 3d;
DEFINE FIELD name ON TABLE tag TYPE string;
DEFINE FIELD color ON TABLE tag TYPE option<string>;

--

LET $organization = (CREATE organization SET name = 'Test Org 1', description = 'Testing organization one.', status = 'pending');
LET $component1 = (CREATE component SET name = 'parent', description = 'b', status = 'c');
LET $component2 = (CREATE component SET name = 'child', description = 'b', status = 'c');
LET $component3 = (CREATE component SET name = 'child', description = 'ccc', status = 'cd');
RELATE $organization->organization_component->$component1;
RELATE $component2->component_link->$component1;
RELATE $component3->component_link->$component2;

LET $property = (CREATE property SET name = 'x', description = 'y');
RELATE $component1->component_property->$property;

LET $tag = (CREATE tag SET name = 'tag1');
RELATE $tag->tag_component->$component1;