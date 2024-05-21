import os
import xml.etree.ElementTree as ET
import requests
import urllib3
import re

# Suppress only the single InsecureRequestWarning from urllib3 needed for making insecure requests
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

def extract_dependencies(pom_file):
    tree = ET.parse(pom_file)
    root = tree.getroot()

    # Define the namespace for Maven POM XML
    ns = {'m': 'http://maven.apache.org/POM/4.0.0'}

    # Collect properties to resolve placeholders
    properties = {}
    properties_element = root.find('m:properties', ns)
    if properties_element is not None:
        for prop in properties_element:
            properties[prop.tag.split('}', 1)[1]] = prop.text

    dependencies = set()

    # Find all dependency elements in the pom.xml
    for dependency in root.findall('m:dependencies/m:dependency', ns):
        group_id = dependency.find('m:groupId', ns).text
        artifact_id = dependency.find('m:artifactId', ns).text
        version = dependency.find('m:version', ns).text if dependency.find('m:version', ns) is not None else 'No version specified'

        # Resolve version placeholders
        version = resolve_version(version, properties)

        # Fetch license information if available
        if version == 'No version specified':
            latest_version = fetch_latest_version(group_id, artifact_id)
            if latest_version:
                version = latest_version
        license_info = fetch_license_info(group_id, artifact_id, version)

        # Add dependency as a tuple to the set
        dependencies.add((group_id, artifact_id, version, ', '.join(license_info)))

    return dependencies

def resolve_version(version, properties):
    pattern = re.compile(r'\$\{(.+?)\}')
    match = pattern.match(version)
    if match:
        var_name = match.group(1)
        return properties.get(var_name, f"Unresolved variable: {var_name}")
    return version

def fetch_latest_version(group_id, artifact_id):
    try:
        # Construct the URL to fetch the latest version metadata
        url = f"https://repo1.maven.org/maven2/{group_id.replace('.', '/')}/{artifact_id}/maven-metadata.xml"
        response = requests.get(url, timeout=10, verify=False)
        
        if response.status_code == 200:
            metadata = ET.fromstring(response.text)
            latest = metadata.find('versioning/latest')
            if latest is not None:
                return latest.text
        return None
    except requests.exceptions.RequestException as e:
        print(f"Error fetching latest version: {str(e)}")
        return None

def fetch_license_info(group_id, artifact_id, version):
    try:
        # Construct the URL for Maven Central Repository metadata
        url = f"https://repo1.maven.org/maven2/{group_id.replace('.', '/')}/{artifact_id}/{version}/{artifact_id}-{version}.pom"
        response = requests.get(url, timeout=10, verify=False)

        if response.status_code == 200:
            pom_content = response.text
            pom_tree = ET.ElementTree(ET.fromstring(pom_content))
            pom_root = pom_tree.getroot()

            # Define the namespace for the fetched POM XML
            ns = {'m': 'http://maven.apache.org/POM/4.0.0'}

            licenses = pom_root.findall('m:licenses/m:license/m:name', ns)
            if licenses:
                return [license.text for license in licenses]
        return ["License information not found"]
    except requests.exceptions.RequestException as e:
        return [f"Error fetching license: {str(e)}"]

def process_all_poms(root_dir):
    all_dependencies = set()
    for subdir, _, files in os.walk(root_dir):
        for file in files:
            if file == 'pom.xml':
                pom_path = os.path.join(subdir, file)
                print(f"Processing {pom_path}")
                dependencies = extract_dependencies(pom_path)
                all_dependencies.update(dependencies)
    return all_dependencies

def print_dependencies(dependencies):
    print("Consolidated Dependencies found in all pom.xml files:")
    for dep in sorted(dependencies):
        full_path = f"{dep[0].replace('.', '/')}/{dep[1]}/{dep[2]}"
        print(f"Full Path: {full_path} | Version: {dep[2]} | License: {dep[3]}")

def main():
    root_dir = 'path/to/your/projects'
    all_dependencies = process_all_poms(root_dir)
    print_dependencies(all_dependencies)

if __name__ == "__main__":
    main()
