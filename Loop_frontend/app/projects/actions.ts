'use server'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;

export async function getAllProjects(refresh_token: string) {
  try {
    const response = await fetch(`${API_BASE_URL}/project/get_projects`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${refresh_token}`,
        'Content-Type': 'application/json'
      },
      credentials: 'include',
      mode: 'cors'
    });

    if (response.status === 401) {
      throw new Error('Unauthorized');
    }
    else if (!response.ok) {
      console.log(response)
      throw new Error('Failed to fetch projects');
    }
    return await response.json();
  } catch (error) {
    console.error('Error fetching projects:', error);
    throw error;
  }
}

// // Fetch projects by status
// export async function getProjectsByStatus(status: string) {
//   try {
//     const response = await fetch(`${API_BASE_URL}/api/projects/status?status=${status}`);
//     if (!response.ok) {
//       throw new Error(`Failed to fetch projects with status ${status}`);
//     }
//     return await response.json();
//   } catch (error) {
//     console.error(`Error fetching projects with status ${status}:`, error);
//     throw error;
//   }
// }
